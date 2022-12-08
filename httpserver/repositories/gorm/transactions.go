package gorm

import (
	"context"
	"time"

	"github.com/storyofhis/toko-belanja/httpserver/repositories"
	"github.com/storyofhis/toko-belanja/httpserver/repositories/models"
	"gorm.io/gorm"
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

func (repo *transactionsRepo) GetMyTransaction(ctx context.Context) ([]models.TransactionHistory, error) {
	var transaction []models.TransactionHistory
	return transaction, repo.db.WithContext(ctx).Find(&transaction).Error
}

func (repo *transactionsRepo) GetUserTransaction(ctx context.Context) ([]models.TransactionHistory, error) {
	var transaction []models.TransactionHistory
	return transaction, repo.db.WithContext(ctx).Find(&transaction).Error
}
