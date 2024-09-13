package service

import (
	"context"

	"github.com/aldisypu/go-simple-pos/internal/model/web"
)

type SaleService interface {
	Create(ctx context.Context, request *web.CreateSaleRequest) (*web.SaleResponse, error)
	Delete(ctx context.Context, request *web.DeleteSaleRequest) error
	Get(ctx context.Context, request *web.GetSaleRequest) (*web.SaleResponse, error)
	List(ctx context.Context) ([]web.SaleResponse, error)
}
