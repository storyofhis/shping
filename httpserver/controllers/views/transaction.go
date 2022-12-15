package views

import "time"

type CreateTransaction struct {
	TotalPrice   uint   `json:"total_price"`
	Quantity     uint   `json:"quantity"`
	ProductTitle string `json:"product_title"`
}

type GetMyTransaction struct {
	Id         uint               `json:"id"`
	ProductId  uint               `json:"product_id"`
	UserId     uint               `json:"user_id"`
	Quantity   uint               `json:"quantity"`
	TotalPrice uint               `json:"total_price"`
	Product    ProductTransaction `json:"product"`
}

type GetUserTransaction struct {
	Id         uint               `json:"id"`
	ProductId  uint               `json:"product_id"`
	UserId     uint               `json:"user_id"`
	Quantity   int                `json:"quantity"`
	TotalPrice int                `json:"total_price"`
	Product    ProductTransaction `json:"product"`
	User       UserTransaction    `json:"user"`
}

type ProductTransaction struct {
	Id         uint      `json:"id"`
	Title      string    `json:"title"`
	Price      uint      `json:"price"`
	Stock      int       `json:"stock"`
	CategoryId uint      `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
type UserTransaction struct {
	Id        uint      `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Balance   uint      `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
