package service

import (
	"context"

	"github.com/aldisypu/go-simple-pos/internal/model/converter"
	"github.com/aldisypu/go-simple-pos/internal/model/domain"
	"github.com/aldisypu/go-simple-pos/internal/model/web"
	"github.com/aldisypu/go-simple-pos/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductServiceImpl struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	ProductRepository repository.ProductRepository
}

func NewProductService(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, productRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{
		DB:                db,
		Log:               logger,
		Validate:          validate,
		ProductRepository: productRepository,
	}
}

func (s *ProductServiceImpl) Create(ctx context.Context, request *web.CreateProductRequest) (*web.ProductResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validating request body")
		return nil, fiber.ErrBadRequest
	}

	product := &domain.Product{
		ID:          uuid.NewString(),
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		CategoryId:  request.CategoryId,
	}

	if err := s.ProductRepository.Create(tx, product); err != nil {
		s.Log.WithError(err).Error("failed to creating product")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ProductToResponse(product), nil
}

func (s *ProductServiceImpl) Update(ctx context.Context, request *web.UpdateProductRequest) (*web.ProductResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	product := new(domain.Product)
	if err := s.ProductRepository.FindById(tx, product, request.ID); err != nil {
		s.Log.WithError(err).Error("failed to getting product")
		return nil, fiber.ErrNotFound
	}

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validating request body")
		return nil, fiber.ErrBadRequest
	}

	product.Name = request.Name
	product.Description = request.Description
	product.Price = request.Price
	product.Stock = request.Stock
	product.CategoryId = request.CategoryId

	if err := s.ProductRepository.Update(tx, product); err != nil {
		s.Log.WithError(err).Error("failed to updating product")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ProductToResponse(product), nil
}

func (s *ProductServiceImpl) Delete(ctx context.Context, request *web.DeleteProductRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validating request body")
		return fiber.ErrBadRequest
	}

	product := new(domain.Product)
	if err := s.ProductRepository.FindById(tx, product, request.ID); err != nil {
		s.Log.WithError(err).Error("failed to getting product")
		return fiber.ErrNotFound
	}

	if err := s.ProductRepository.Delete(tx, product); err != nil {
		s.Log.WithError(err).Error("failed to deleting product")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *ProductServiceImpl) Get(ctx context.Context, request *web.GetProductRequest) (*web.ProductResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validating request body")
		return nil, fiber.ErrBadRequest
	}

	product := new(domain.Product)
	if err := s.ProductRepository.FindById(tx, product, request.ID); err != nil {
		s.Log.WithError(err).Error("failed to getting product")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.ProductToResponse(product), nil
}

func (s *ProductServiceImpl) List(ctx context.Context) ([]web.ProductResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	products, err := s.ProductRepository.FindAll(tx)
	if err != nil {
		s.Log.WithError(err).Error("failed to find products")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]web.ProductResponse, len(products))
	for i, product := range products {
		responses[i] = *converter.ProductToResponse(&product)
	}

	return responses, nil
}
