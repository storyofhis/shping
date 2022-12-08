package services

import (
	"context"

	"github.com/storyofhis/toko-belanja/httpserver/controllers/params"
	"github.com/storyofhis/toko-belanja/httpserver/controllers/views"
)

type UserSvc interface {
	Register(ctx context.Context, user *params.Register) *views.Response
	Login(ctx context.Context, user *params.Login) *views.Response
	TopUp(ctx context.Context, id uint, params *params.TopUp) *views.Response
}

type CategorySvc interface {
	CreateCategory(ctx context.Context, category *params.CreateCategory) *views.Response
	// GetCategory(ctx context.Context) *views.Response
	// UpdateCategory(ctx context.Context, category *params.UpdateCategory, id uint) *views.Response
	// DeleteCategory(ctx context.Context, id uint) *views.CreateCategory
}

type ProductSvc interface {
	CreateProduct(ctx context.Context, product *params.CreateProduct) *views.Response
	GetProducts(ctx context.Context) *views.Response
	UpdateProduct(ctx context.Context, product *params.UpdateProduct, productId uint) *views.Response
	DeleteProduct(ctx context.Context, productId uint) *views.Response
}

type TransactionSvc interface {
	CreateTransaction(ctx context.Context, transaction *params.CreateTransactions) *views.Response
	GetMyTransaction(ctx context.Context) *views.Response
}
