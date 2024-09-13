package repository

import (
	"github.com/aldisypu/go-simple-pos/internal/model/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	Log *logrus.Logger
}

func NewProductRepository(log *logrus.Logger) ProductRepository {
	return &ProductRepositoryImpl{
		Log: log,
	}
}

func (r *ProductRepositoryImpl) Create(db *gorm.DB, product *domain.Product) error {
	return db.Create(product).Error
}

func (r *ProductRepositoryImpl) Update(db *gorm.DB, product *domain.Product) error {
	return db.Save(product).Error
}

func (r *ProductRepositoryImpl) Delete(db *gorm.DB, product *domain.Product) error {
	return db.Delete(product).Error
}

func (r *ProductRepositoryImpl) FindById(db *gorm.DB, product *domain.Product, id string) error {
	return db.Where("id = ?", id).Take(product).Error
}

func (r *ProductRepositoryImpl) FindAll(db *gorm.DB) ([]domain.Product, error) {
	var products []domain.Product
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepositoryImpl) GetPriceById(db *gorm.DB, productId string) (float64, error) {
	var product domain.Product
	if err := db.Where("id = ?", productId).Select("price").First(&product).Error; err != nil {
		return 0, err
	}
	return product.Price, nil
}

func (r *ProductRepositoryImpl) IncreaseStock(db *gorm.DB, productId string, quantity int) error {
	return db.Model(&domain.Product{}).
		Where("id = ?", productId).
		UpdateColumn("stock", gorm.Expr("stock + ?", quantity)).Error
}

func (r *ProductRepositoryImpl) DecreaseStock(db *gorm.DB, productId string, quantity int) error {
	return db.Model(&domain.Product{}).
		Where("id = ?", productId).
		UpdateColumn("stock", gorm.Expr("stock - ?", quantity)).Error
}
