package web

type ProductResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CategoryId  string  `json:"category_id"`
	CreatedAt   int64   `json:"created_at"`
	UpdatedAt   int64   `json:"updated_at"`
}

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required,max=100"`
	Description string  `json:"description" validate:"required,max=255"`
	Price       float64 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	CategoryId  string  `json:"category_id" validate:"required"`
}

type UpdateProductRequest struct {
	ID          string  `json:"-" validate:"required,max=100,uuid"`
	Name        string  `json:"name" validate:"required,max=100"`
	Description string  `json:"description" validate:"required,max=255"`
	Price       float64 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	CategoryId  string  `json:"category_id" validate:"required"`
}

type DeleteProductRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type GetProductRequest struct {
	ID         string `json:"-" validate:"required,max=100,uuid"`
	CategoryId string `json:"-" validate:"required,max=100,uuid"`
}
