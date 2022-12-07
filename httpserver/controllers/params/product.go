package params

type CreateProduct struct {
	Title      string `json:"title" validate:"required"`
	Price      uint   `json:"price" validate:"required,min=0,max=50000000"`
	Stock      int    `json:"stock" validate:"required,min=5"`
	CategoryId uint   `json:"category_id" validate:"required"`
}

type UpdateProduct struct {
	Title      string `json:"title" validate:"required"`
	Price      uint   `json:"price" validate:"required,min=0,max=50000000"`
	Stock      int    `json:"stock" validate:"required,min=5"`
	CategoryId uint   `json:"category_id" validate:"required"`
}
