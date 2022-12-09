package gorm

import (
	"context"
	"time"

	"github.com/storyofhis/toko-belanja/httpserver/repositories"
	"github.com/storyofhis/toko-belanja/httpserver/repositories/models"
	"gorm.io/gorm"
)

type productsRepo struct {
	db *gorm.DB
}

func NewProductsRepo(db *gorm.DB) repositories.ProductsRepo {
	return &productsRepo{
		db: db,
	}
}

func (repo *productsRepo) CreateProduct(ctx context.Context, product *models.Products) error {
	product.CreatedAt = time.Now()
	return repo.db.WithContext(ctx).Create(product).Error
}

func (repo *productsRepo) GetAllProducts(ctx context.Context) ([]models.Products, error) {
	var products []models.Products

	if err := repo.db.WithContext(ctx).Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}

func (repo *productsRepo) GetProductById(ctx context.Context, id uint) (*models.Products, error) {
	products := new(models.Products)
	err := repo.db.WithContext(ctx).Where("id = ?", id).Take(products).Error
	return products, err
}

func (repo *productsRepo) UpdateProduct(ctx context.Context, product *models.Products, id uint) error {
	product.UpdatedAt = time.Now()
	return repo.db.WithContext(ctx).Model(product).Where("id = ?", id).Updates(*product).Error
}

func (repo *productsRepo) DeleteProduct(ctx context.Context, id uint) error {
	return repo.db.WithContext(ctx).Delete(&models.Products{}, "id = ?", id).Error
}

func (repo *productsRepo) FindProductsByCategoryId(ctx context.Context, id uint) ([]models.Products, error) {
	var products []models.Products

	if err := repo.db.WithContext(ctx).Where("category_id = ?", id).Find(&products).Error; err != nil {
		return products, err
	}
	return products, nil
}
