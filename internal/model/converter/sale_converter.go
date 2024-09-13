package converter

import (
	"github.com/aldisypu/go-simple-pos/internal/model/domain"
	"github.com/aldisypu/go-simple-pos/internal/model/web"
)

func SaleToResponse(sale *domain.Sale) *web.SaleResponse {
	return &web.SaleResponse{
		ID:        sale.ID,
		SaleDate:  sale.SaleDate,
		Total:     sale.Total,
		CreatedAt: sale.CreatedAt,
		UpdatedAt: sale.UpdatedAt,
	}
}
