package gorm

import (
	"context"
	"github.com/storyofhis/toko-belanja/httpserver/repositories"
	"github.com/storyofhis/toko-belanja/httpserver/repositories/models"
	"gorm.io/gorm"
	"time"
)

type transactionsRepo struct {
	db *gorm.DB
}

func NewTransactionsRepo(db *gorm.DB) repositories.TransactionsRepo {
	return &transactionsRepo{
		db: db,
	}
}

func (repo *transactionsRepo) CreateTransaction(ctx context.Context, transaction *models.TransactionHistory) error {
	transaction.CreatedAt = time.Now()
	return repo.db.WithContext(ctx).Create(transaction).Error
}
