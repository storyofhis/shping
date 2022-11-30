package views

import "time"

type CreateProduct struct {
	Id         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      uint      `json:"price"`
	Stock      int       `json:"stock"`
	CategoryId uint      `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type GetProduct struct {
	Id         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      uint      `json:"price"`
	Stock      int       `json:"stock"`
	CategoryId uint      `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type UpdateProduct struct {
	Id         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      uint      `json:"price"`
	Stock      int       `json:"stock"`
	CategoryId uint      `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type DeleteProduct struct {
	Message string `json:"message"`
}
