package repository

import (
	"github.com/aldisypu/go-simple-pos/internal/model/domain"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(db *gorm.DB, category *domain.Category) error
	Update(db *gorm.DB, category *domain.Category) error
	Delete(db *gorm.DB, category *domain.Category) error
	FindById(db *gorm.DB, category *domain.Category, id string) error
	FindAll(db *gorm.DB) ([]domain.Category, error)
}
