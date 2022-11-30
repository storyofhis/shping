package gorm

import (
	"context"
	"log"
	"time"

	"github.com/storyofhis/toko-belanja/httpserver/repositories"
	"github.com/storyofhis/toko-belanja/httpserver/repositories/models"
	"gorm.io/gorm"
)

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) repositories.CategoryRepo {
	return &categoryRepo{
		db: db,
	}
}

func (repo *categoryRepo) CreateCategory(ctx context.Context, category *models.Categories) error {
	category.CreatedAt = time.Now()
	return repo.db.WithContext(ctx).Create(category).Error
}

func (repo *categoryRepo) GetCategory(ctx context.Context) ([]models.Categories, error) {
	var categories []models.Categories

	err := repo.db.WithContext(ctx).Find(&categories).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return categories, err
}

func (repo *categoryRepo) FindCategoryById(ctx context.Context, id uint) (*models.Categories, error) {
	category := new(models.Categories)
	err := repo.db.WithContext(ctx).Where("id = ?", id).Take(category).Error
	return category, err
}

func (repo *categoryRepo) UpdateCategory(ctx context.Context, category *models.Categories, id uint) error {
	category.UpdatedAt = time.Now()
	return repo.db.WithContext(ctx).Model(category).Where("id = ?", id).Updates(*category).Error
}

func (repo *categoryRepo) DeleteCategory(ctx context.Context, id uint) error {
	return repo.db.WithContext(ctx).Delete(&models.Categories{}, "id = ?", id).Error
}
