package converter

import (
	"github.com/aldisypu/go-simple-pos/internal/model/domain"
	"github.com/aldisypu/go-simple-pos/internal/model/web"
)

func SalesDetailToResponse(salesDetail *domain.SalesDetail) *web.SalesDetailResponse {
	return &web.SalesDetailResponse{
		ID:        salesDetail.ID,
		SaleId:    salesDetail.SaleId,
		ProductId: salesDetail.ProductId,
		Quantity:  salesDetail.Quantity,
		Price:     salesDetail.Price,
		CreatedAt: salesDetail.CreatedAt,
		UpdatedAt: salesDetail.UpdatedAt,
	}
}
