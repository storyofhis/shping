package repositories

import (
	"context"

	"github.com/storyofhis/toko-belanja/httpserver/repositories/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.Users) error
	FindUserByEmail(ctx context.Context, email string) (*models.Users, error)
	FindUserById(ctx context.Context, id uint) (*models.Users, error)
	UpdateUser(ctx context.Context, user *models.Users) error
}

type CategoryRepo interface {
	CreateCategory(ctx context.Context, category *models.Categories) error
	GetCategories(ctx context.Context) ([]models.Categories, error)
	FindCategoryById(ctx context.Context, id uint) (*models.Categories, error)
	UpdateCategory(ctx context.Context, category *models.Categories, id uint) error
	DeleteCategory(ctx context.Context, id uint) error
}

type ProductsRepo interface {
	CreateProduct(ctx context.Context, product *models.Products) error
	GetAllProducts(ctx context.Context) ([]models.Products, error)
	UpdateProduct(ctx context.Context, product *models.Products, id uint) error
	GetProductById(ctx context.Context, id uint) (*models.Products, error)
	DeleteProduct(ctx context.Context, id uint) error
	FindProductsByCategoryId(ctx context.Context, id uint) ([]models.Products, error)
}

type TransactionsRepo interface {
	CreateTransaction(ctx context.Context, transaction *models.TransactionHistory) error
	GetMyTransaction(ctx context.Context, id uint) ([]models.TransactionHistory, error)
	GetUserTransaction(ctx context.Context) ([]models.TransactionHistory, error)
}
