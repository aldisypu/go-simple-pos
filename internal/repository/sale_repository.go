package repository

import (
	"github.com/aldisypu/go-simple-pos/internal/model/domain"
	"gorm.io/gorm"
)

type SaleRepository interface {
	Create(tx *gorm.DB, sale *domain.Sale) error
	Update(db *gorm.DB, sale *domain.Sale) error
	Delete(db *gorm.DB, sale *domain.Sale) error
	FindById(db *gorm.DB, sale *domain.Sale, id string) error
	FindAll(db *gorm.DB) ([]domain.Sale, error)
}
