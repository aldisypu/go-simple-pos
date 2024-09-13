package repository

import (
	"github.com/aldisypu/go-simple-pos/internal/model/domain"
	"gorm.io/gorm"
)

type SalesDetailRepository interface {
	Update(db *gorm.DB, salesDetail *domain.SalesDetail) error
	Delete(db *gorm.DB, salesDetail *domain.SalesDetail) error
	DeleteBySaleId(tx *gorm.DB, saleId string) error
	FindByIdAndSaleId(db *gorm.DB, salesDetail *domain.SalesDetail, id string, saleId string) error
	FindAllBySaleId(tx *gorm.DB, saleId string) ([]domain.SalesDetail, error)
}
