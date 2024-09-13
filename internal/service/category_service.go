package service

import (
	"context"

	"github.com/aldisypu/go-simple-pos/internal/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request *web.CreateCategoryRequest) (*web.CategoryResponse, error)
	Update(ctx context.Context, request *web.UpdateCategoryRequest) (*web.CategoryResponse, error)
	Delete(ctx context.Context, request *web.DeleteCategoryRequest) error
	Get(ctx context.Context, request *web.GetCategoryRequest) (*web.CategoryResponse, error)
	List(ctx context.Context) ([]web.CategoryResponse, error)
}
