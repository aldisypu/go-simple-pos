package repository

import (
	"github.com/aldisypu/go-simple-pos/internal/model/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	Log *logrus.Logger
}

func NewCategoryRepository(log *logrus.Logger) CategoryRepository {
	return &CategoryRepositoryImpl{
		Log: log,
	}
}

func (r *CategoryRepositoryImpl) Create(db *gorm.DB, category *domain.Category) error {
	return db.Create(category).Error
}

func (r *CategoryRepositoryImpl) Update(db *gorm.DB, category *domain.Category) error {
	return db.Save(category).Error
}

func (r *CategoryRepositoryImpl) Delete(db *gorm.DB, category *domain.Category) error {
	return db.Delete(category).Error
}

func (r *CategoryRepositoryImpl) FindById(db *gorm.DB, category *domain.Category, id string) error {
	return db.Where("id = ?", id).Take(category).Error
}

func (r *CategoryRepositoryImpl) FindAll(db *gorm.DB) ([]domain.Category, error) {
	var categories []domain.Category
	if err := db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}
