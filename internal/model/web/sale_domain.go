package web

type SaleResponse struct {
	ID        string                `json:"id"`
	SaleDate  int64                 `json:"sale_date"`
	Total     float64               `json:"total"`
	CreatedAt int64                 `json:"created_at"`
	UpdatedAt int64                 `json:"updated_at"`
	Details   []SalesDetailResponse `json:"details,omitempty"`
}

type CreateSaleRequest struct {
	Details []CreateSalesDetailRequest `json:"details"`
}

type DeleteSaleRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type GetSaleRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}
