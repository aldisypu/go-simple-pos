package service

import (
	"context"

	"github.com/aldisypu/go-simple-pos/internal/model/converter"
	"github.com/aldisypu/go-simple-pos/internal/model/domain"
	"github.com/aldisypu/go-simple-pos/internal/model/web"
	"github.com/aldisypu/go-simple-pos/internal/repository"
	"github.com/google/uuid"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SaleServiceImpl struct {
	DB                    *gorm.DB
	Log                   *logrus.Logger
	Validate              *validator.Validate
	SaleRepository        repository.SaleRepository
	ProductRepository     repository.ProductRepository
	SalesDetailRepository repository.SalesDetailRepository
}

func NewSaleService(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, saleRepository repository.SaleRepository, productRepository repository.ProductRepository, salesDetailRepository repository.SalesDetailRepository) SaleService {
	return &SaleServiceImpl{
		DB:                    db,
		Log:                   logger,
		Validate:              validate,
		SaleRepository:        saleRepository,
		ProductRepository:     productRepository,
		SalesDetailRepository: salesDetailRepository,
	}
}

func (s *SaleServiceImpl) Create(ctx context.Context, request *web.CreateSaleRequest) (*web.SaleResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validating request body")
		return nil, fiber.ErrBadRequest
	}

	sale := &domain.Sale{
		ID: uuid.NewString(),
	}

	var total float64
	for _, detailRequest := range request.Details {
		price, err := s.ProductRepository.GetPriceById(tx, detailRequest.ProductId)
		if err != nil {
			s.Log.WithError(err).Error("failed to fetching product price")
			return nil, fiber.ErrInternalServerError
		}

		if err := s.ProductRepository.DecreaseStock(tx, detailRequest.ProductId, detailRequest.Quantity); err != nil {
			s.Log.WithError(err).Error("failed to reducing product stock")
			return nil, fiber.ErrInternalServerError
		}

		detail := domain.SalesDetail{
			ID:        uuid.NewString(),
			SaleId:    sale.ID,
			ProductId: detailRequest.ProductId,
			Quantity:  detailRequest.Quantity,
			Price:     price,
		}

		sale.Details = append(sale.Details, detail)
		total += detail.Price * float64(detail.Quantity)
	}

	sale.Total = total

	if err := s.SaleRepository.Create(tx, sale); err != nil {
		s.Log.WithError(err).Error("failed to creating sale")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.SaleToResponse(sale), nil
}

func (s *SaleServiceImpl) Delete(ctx context.Context, request *web.DeleteSaleRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validating request body")
		return fiber.ErrBadRequest
	}

	sale := new(domain.Sale)
	if err := s.SaleRepository.FindById(tx, sale, request.ID); err != nil {
		s.Log.WithError(err).Error("failed to getting sale")
		return fiber.ErrNotFound
	}

	if err := s.SalesDetailRepository.DeleteBySaleId(tx, sale.ID); err != nil {
		s.Log.WithError(err).Error("failed to deleting sale details")
		return fiber.ErrNotFound
	}

	if err := s.SaleRepository.Delete(tx, sale); err != nil {
		s.Log.WithError(err).Error("failed to deleting sale")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *SaleServiceImpl) Get(ctx context.Context, request *web.GetSaleRequest) (*web.SaleResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validating request body")
		return nil, fiber.ErrBadRequest
	}

	sale := new(domain.Sale)
	if err := s.SaleRepository.FindById(tx, sale, request.ID); err != nil {
		s.Log.WithError(err).Error("failed to getting sale")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.SaleToResponse(sale), nil
}

func (s *SaleServiceImpl) List(ctx context.Context) ([]web.SaleResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	sales, err := s.SaleRepository.FindAll(tx)
	if err != nil {
		s.Log.WithError(err).Error("failed to find sales")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]web.SaleResponse, len(sales))
	for i, sale := range sales {
		responses[i] = *converter.SaleToResponse(&sale)
	}

	return responses, nil
}
