package params

type CreateProduct struct {
	Title      string `json:"title" validate:"required"`
	Price      uint   `json:"price" validate:"required"`
	Stock      int    `json:"stock" validate:"required"`
	CategoryId uint   `json:"category_id" validate:"required"`
}

type UpdateProduct struct {
	Title      string `json:"title" validate:"required"`
	Price      uint   `json:"price" validate:"required"`
	Stock      int    `json:"stock" validate:"required"`
	CategoryId uint   `json:"category_id" validate:"required"`
}
