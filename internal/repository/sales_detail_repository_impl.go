package repository

import (
	"github.com/aldisypu/go-simple-pos/internal/model/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SalesDetailRepositoryImpl struct {
	Log *logrus.Logger
}

func NewSalesDetailRepository(log *logrus.Logger) SalesDetailRepository {
	return &SalesDetailRepositoryImpl{
		Log: log,
	}
}

func (r *SalesDetailRepositoryImpl) Update(db *gorm.DB, salesDetail *domain.SalesDetail) error {
	return db.Model(&salesDetail).Update("quantity", salesDetail.Quantity).Error
}

func (r *SalesDetailRepositoryImpl) Delete(db *gorm.DB, salesDetail *domain.SalesDetail) error {
	return db.Delete(salesDetail).Error
}

func (r *SalesDetailRepositoryImpl) DeleteBySaleId(tx *gorm.DB, saleId string) error {
	return tx.Where("sale_id = ?", saleId).Delete(&domain.SalesDetail{}).Error
}

func (r *SalesDetailRepositoryImpl) FindByIdAndSaleId(db *gorm.DB, salesDetail *domain.SalesDetail, id string, saleId string) error {
	return db.Where("id = ? AND sale_id = ?", id, saleId).First(salesDetail).Error
}

func (r *SalesDetailRepositoryImpl) FindAllBySaleId(tx *gorm.DB, saleId string) ([]domain.SalesDetail, error) {
	var salesDetails []domain.SalesDetail
	if err := tx.Where("sale_id = ?", saleId).Find(&salesDetails).Error; err != nil {
		return nil, err
	}

	return salesDetails, nil
}
