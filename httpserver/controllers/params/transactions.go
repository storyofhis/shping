package params

type CreateTransactions struct {
	ProductId uint `json:"product_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required"`
}
