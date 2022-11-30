package views

import "time"

type CreateCategory struct {
	Id                uint      `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount uint      `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
}

type GetCategory struct {
	Id                uint      `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount uint      `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type UpdateCategory struct {
	Id                uint      `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount uint      `json:"sold_product_amount"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type DeleteCategory struct {
	Message string `json:"message"`
}
