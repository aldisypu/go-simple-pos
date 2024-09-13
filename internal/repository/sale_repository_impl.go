package repository

import (
	"github.com/aldisypu/go-simple-pos/internal/model/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SaleRepositoryImpl struct {
	Log *logrus.Logger
}

func NewSaleRepository(log *logrus.Logger) SaleRepository {
	return &SaleRepositoryImpl{
		Log: log,
	}
}

func (r *SaleRepositoryImpl) Create(tx *gorm.DB, sale *domain.Sale) error {
	return tx.Save(sale).Error
}

func (r *SaleRepositoryImpl) Update(db *gorm.DB, sale *domain.Sale) error {
	return db.Model(&sale).Update("total", sale.Total).Error
}

func (r *SaleRepositoryImpl) Delete(db *gorm.DB, sale *domain.Sale) error {
	return db.Delete(sale).Error
}

func (r *SaleRepositoryImpl) FindById(db *gorm.DB, sale *domain.Sale, id string) error {
	return db.Where("id = ?", id).Take(sale).Error
}

func (r *SaleRepositoryImpl) FindAll(db *gorm.DB) ([]domain.Sale, error) {
	var sales []domain.Sale
	if err := db.Find(&sales).Error; err != nil {
		return nil, err
	}

	return sales, nil
}
