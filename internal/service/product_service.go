package service

import (
	"context"

	"github.com/aldisypu/go-simple-pos/internal/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request *web.CreateProductRequest) (*web.ProductResponse, error)
	Update(ctx context.Context, request *web.UpdateProductRequest) (*web.ProductResponse, error)
	Delete(ctx context.Context, request *web.DeleteProductRequest) error
	Get(ctx context.Context, request *web.GetProductRequest) (*web.ProductResponse, error)
	List(ctx context.Context) ([]web.ProductResponse, error)
}
