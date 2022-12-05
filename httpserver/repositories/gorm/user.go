package gorm

import (
	"context"
	"strings"
	"time"

	"github.com/storyofhis/toko-belanja/httpserver/repositories"
	"github.com/storyofhis/toko-belanja/httpserver/repositories/models"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repositories.UserRepo {
	return &userRepo{
		db: db,
	}
}

func (repo *userRepo) CreateUser(ctx context.Context, user *models.Users) error {
	user.CreatedAt = time.Now()
	return repo.db.WithContext(ctx).Create(user).Error
}

func (repo *userRepo) FindUserByEmail(ctx context.Context, email string) (*models.Users, error) {
	user := new(models.Users)
	return user, repo.db.WithContext(ctx).Where("LOWER(email) = ?", strings.ToLower(email)).Take(user).Error
}

func (repo *userRepo) FindUserById(ctx context.Context, id uint) (*models.Users, error) {
	user := new(models.Users)
	return user, repo.db.WithContext(ctx).Where("id = ?", id).Take(user).Error
}

func (repo *userRepo) UpdateUser(ctx context.Context, user *models.Users) error {
	user.UpdatedAt = time.Now()
	return repo.db.WithContext(ctx).Model(user).Updates(*user).Error
}
