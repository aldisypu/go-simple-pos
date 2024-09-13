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

type CategoryServiceImpl struct {
	DB                 *gorm.DB
	Log                *logrus.Logger
	Validate           *validator.Validate
	CategoryRepository repository.CategoryRepository
}

func NewCategoryService(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, categoryRepository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		DB:                 db,
		Log:                logger,
		Validate:           validate,
		CategoryRepository: categoryRepository,
	}
}

func (s *CategoryServiceImpl) Create(ctx context.Context, request *web.CreateCategoryRequest) (*web.CategoryResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validating request body")
		return nil, fiber.ErrBadRequest
	}

	category := &domain.Category{
		ID:   uuid.NewString(),
		Name: request.Name,
	}

	if err := s.CategoryRepository.Create(tx, category); err != nil {
		s.Log.WithError(err).Error("failed to creating category")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CategoryToResponse(category), nil
}

func (s *CategoryServiceImpl) Update(ctx context.Context, request *web.UpdateCategoryRequest) (*web.CategoryResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	category := new(domain.Category)
	if err := s.CategoryRepository.FindById(tx, category, request.ID); err != nil {
		s.Log.WithError(err).Error("failed to getting category")
		return nil, fiber.ErrNotFound
	}

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validating request body")
		return nil, fiber.ErrBadRequest
	}

	category.Name = request.Name

	if err := s.CategoryRepository.Update(tx, category); err != nil {
		s.Log.WithError(err).Error("failed to updating category")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CategoryToResponse(category), nil
}

func (s *CategoryServiceImpl) Delete(ctx context.Context, request *web.DeleteCategoryRequest) error {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validating request body")
		return fiber.ErrBadRequest
	}

	category := new(domain.Category)
	if err := s.CategoryRepository.FindById(tx, category, request.ID); err != nil {
		s.Log.WithError(err).Error("failed to getting category")
		return fiber.ErrNotFound
	}

	if err := s.CategoryRepository.Delete(tx, category); err != nil {
		s.Log.WithError(err).Error("failed to deleting category")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (s *CategoryServiceImpl) Get(ctx context.Context, request *web.GetCategoryRequest) (*web.CategoryResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := s.Validate.Struct(request); err != nil {
		s.Log.WithError(err).Error("failed to validating request body")
		return nil, fiber.ErrBadRequest
	}

	category := new(domain.Category)
	if err := s.CategoryRepository.FindById(tx, category, request.ID); err != nil {
		s.Log.WithError(err).Error("failed to getting category")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CategoryToResponse(category), nil
}

func (s *CategoryServiceImpl) List(ctx context.Context) ([]web.CategoryResponse, error) {
	tx := s.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	categories, err := s.CategoryRepository.FindAll(tx)
	if err != nil {
		s.Log.WithError(err).Error("failed to find categories")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		s.Log.WithError(err).Error("failed to commit transaction")
		return nil, fiber.ErrInternalServerError
	}

	responses := make([]web.CategoryResponse, len(categories))
	for i, category := range categories {
		responses[i] = *converter.CategoryToResponse(&category)
	}

	return responses, nil
}
