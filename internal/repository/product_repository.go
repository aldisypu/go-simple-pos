package repository

import (
	"github.com/aldisypu/go-simple-pos/internal/model/domain"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(db *gorm.DB, product *domain.Product) error
	Update(db *gorm.DB, product *domain.Product) error
	Delete(db *gorm.DB, product *domain.Product) error
	FindById(db *gorm.DB, product *domain.Product, id string) error
	FindAll(db *gorm.DB) ([]domain.Product, error)
	GetPriceById(db *gorm.DB, productId string) (float64, error)
	IncreaseStock(db *gorm.DB, productId string, quantity int) error
	DecreaseStock(db *gorm.DB, productId string, quantity int) error
}
