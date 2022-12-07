package views

type CreateTransaction struct {
	TotalPrice   uint   `json:"total_price"`
	Quantity     uint   `json:"quantity"`
	ProductTitle string `json:"product_title"`
}
