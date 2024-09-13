package service

import (
	"context"

	"github.com/aldisypu/go-simple-pos/internal/model/web"
)

type SalesDetailService interface {
	Update(ctx context.Context, request *web.UpdateSalesDetailRequest) (*web.SalesDetailResponse, error)
	Delete(ctx context.Context, request *web.DeleteSalesDetailRequest) error
	Get(ctx context.Context, request *web.GetSalesDetailRequest) (*web.SalesDetailResponse, error)
	List(ctx context.Context, request *web.ListSalesDetailRequest) ([]web.SalesDetailResponse, error)
}
