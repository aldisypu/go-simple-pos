package service

import (
	"context"

	"github.com/aldisypu/go-simple-pos/internal/model/converter"
	"github.com/aldisypu/go-simple-pos/internal/model/domain"
	"github.com/aldisypu/go-simple-pos/internal/model/web"
	"github.com/aldisypu/go-simple-pos/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SalesDetailServiceImpl struct {
	DB                    *gorm.DB
	Log                   *logrus.Logger
	Validate              *validator.Validate
	ProductRepository     repository.ProductRepository
	SaleRepository        repository.SaleRepository
	SalesDetailRepository repository.SalesDetailRepository
}

func NewSalesDetailService(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, productRepository repository.ProductRepository, salesRepository repository.SaleRepository, salesDetailRepository repository.SalesDetailRepository) SalesDetailService {
	return &SalesDetailServiceImpl{
		DB:                    db,
		Log:                   logger,
		Validate:              validate,
		ProductRepository:     productRepository,
		SaleRepository:        salesRepository,
		SalesDetailRepository: salesDetailRepository,
	}
}

func (s *SalesDetailServiceImpl) Update(ctx context.Context, request *web.UpdateSalesDetailRequest) (*web.SalesDetailResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validating request body")
		return nil, fiber.ErrBadRequest
	}

	salesDetail := new(domain.SalesDetail)
	if err := s.SalesDetailRepository.FindByIdAndSaleId(tx, salesDetail, request.ID, request.SaleId); err != nil {
		s.Log.WithError(err).Error("failed to getting sales detail")
		return nil, fiber.ErrNotFound
	}

	quantityChange := request.Quantity - salesDetail.Quantity

	if quantityChange > 0 {
		if err := s.ProductRepository.DecreaseStock(tx, salesDetail.ProductId, quantityChange); err != nil {
			s.Log.WithError(err).Error("failed to decrease product stock")
			return nil, fiber.ErrInternalServerError
		}
	} else if quantityChange < 0 {
		if err := s.ProductRepository.IncreaseStock(tx, salesDetail.ProductId, -quantityChange); err != nil {
			s.Log.WithError(err).Error("failed to increase product stock")
			return nil, fiber.ErrInternalServerError
		}
	}

	salesDetail.Quantity = request.Quantity
	if err := s.SalesDetailRepository.Update(tx, salesDetail); err != nil {
		s.Log.WithError(err).Error("failed to updating sales detail")
		return nil, fiber.ErrInternalServerError
	}

	sale := new(domain.Sale)
	if err := s.SaleRepository.FindById(tx, sale, request.SaleId); err != nil {
		s.Log.WithError(err).Error("failed to get sale")
		return nil, fiber.ErrNotFound
	}

	totalPriceChange := float64(quantityChange) * salesDetail.Price
	sale.Total += totalPriceChange
	if err := s.SaleRepository.Update(tx, sale); err != nil {
		s.Log.WithError(err).Error("failed to updating sale total")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.SalesDetailToResponse(salesDetail), nil
}

func (s *SalesDetailServiceImpl) Delete(ctx context.Context, request *web.DeleteSalesDetailRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validating request body")
		return fiber.ErrBadRequest
	}

	salesDetail := new(domain.SalesDetail)
	if err := s.SalesDetailRepository.FindByIdAndSaleId(tx, salesDetail, request.ID, request.SaleId); err != nil {
		s.Log.WithError(err).Error("failed to getting sales detail")
		return fiber.ErrNotFound
	}

	totalPriceToSubtract := salesDetail.Price * float64(salesDetail.Quantity)
	quantityToAddBack := salesDetail.Quantity
	productId := salesDetail.ProductId

	if err := s.ProductRepository.IncreaseStock(tx, productId, quantityToAddBack); err != nil {
		s.Log.WithError(err).Error("failed to update product stock")
		return fiber.ErrInternalServerError
	}

	if err := s.SalesDetailRepository.Delete(tx, salesDetail); err != nil {
		s.Log.WithError(err).Error("failed to deleting sales detail")
		return fiber.ErrInternalServerError
	}

	sale := new(domain.Sale)
	if err := s.SaleRepository.FindById(tx, sale, request.SaleId); err != nil {
		s.Log.WithError(err).Error("failed to get sale")
		return fiber.ErrNotFound
	}

	sale.Total -= totalPriceToSubtract
	if err := s.SaleRepository.Update(tx, sale); err != nil {
		s.Log.WithError(err).Error("failed to update sale total")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *SalesDetailServiceImpl) Get(ctx context.Context, request *web.GetSalesDetailRequest) (*web.SalesDetailResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validating request body")
		return nil, fiber.ErrBadRequest
	}

	salesDetail := new(domain.SalesDetail)
	if err := s.SalesDetailRepository.FindByIdAndSaleId(tx, salesDetail, request.ID, request.SaleId); err != nil {
		s.Log.WithError(err).Error("failed to getting sales detail")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.SalesDetailToResponse(salesDetail), nil
}

func (s *SalesDetailServiceImpl) List(ctx context.Context, request *web.ListSalesDetailRequest) ([]web.SalesDetailResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	sale := new(domain.Sale)
	if err := s.SaleRepository.FindById(tx, sale, request.SaleId); err != nil {
		s.Log.WithError(err).Error("failed to find sale")
		return nil, fiber.ErrNotFound
	}

	salesDetails, err := s.SalesDetailRepository.FindAllBySaleId(tx, sale.ID)
	if err != nil {
		s.Log.WithError(err).Error("failed to find sales details")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]web.SalesDetailResponse, len(salesDetails))
	for i, salesDetail := range salesDetails {
		responses[i] = *converter.SalesDetailToResponse(&salesDetail)
	}

	return responses, nil
}
