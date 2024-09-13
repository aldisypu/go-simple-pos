package web

type CategoryResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,max=100"`
}

type UpdateCategoryRequest struct {
	ID   string `json:"-" validate:"required,max=100,uuid"`
	Name string `json:"name" validate:"required,max=100"`
}

type DeleteCategoryRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}

type GetCategoryRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}
