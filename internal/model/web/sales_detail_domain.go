package web

type SalesDetailResponse struct {
	ID        string  `json:"id"`
	SaleId    string  `json:"sale_id"`
	ProductId string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	CreatedAt int64   `json:"created_at"`
	UpdatedAt int64   `json:"updated_at"`
}

type CreateSalesDetailRequest struct {
	SaleId    string  `json:"sale_id" validate:"required"`
	ProductId string  `json:"product_id" validate:"required"`
	Quantity  int     `json:"quantity" validate:"required"`
	Price     float64 `json:"price" validate:"required"`
}

type UpdateSalesDetailRequest struct {
	ID        string `json:"-" validate:"required,max=100,uuid"`
	SaleId    string `json:"-" validate:"required,max=100,uuid"`
	ProductId string `json:"-"`
	Quantity  int    `json:"quantity" validate:"required"`
}

type DeleteSalesDetailRequest struct {
	ID     string `json:"-" validate:"required,max=100,uuid"`
	SaleId string `json:"-" validate:"required,max=100,uuid"`
}

type GetSalesDetailRequest struct {
	ID     string `json:"-" validate:"required,max=100,uuid"`
	SaleId string `json:"-" validate:"required,max=100,uuid"`
}

type ListSalesDetailRequest struct {
	SaleId string `json:"-" validate:"required,max=100,uuid"`
}
