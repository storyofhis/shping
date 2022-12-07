package services

import (
	"context"
	"errors"
	"github.com/storyofhis/toko-belanja/common"
	"github.com/storyofhis/toko-belanja/httpserver/controllers/params"
	"github.com/storyofhis/toko-belanja/httpserver/controllers/views"
	"github.com/storyofhis/toko-belanja/httpserver/repositories"
	"github.com/storyofhis/toko-belanja/httpserver/repositories/models"
	"gorm.io/gorm"
	"net/http"
)

type transactionSvc struct {
	repo         repositories.TransactionsRepo
	productRepo  repositories.ProductsRepo
	userRepo     repositories.UserRepo
	categoryRepo repositories.CategoryRepo
}

func NewTransactionSvc(repo repositories.TransactionsRepo, productRepo repositories.ProductsRepo, userRepo repositories.UserRepo, categoryRepo repositories.CategoryRepo) TransactionSvc {
	return &transactionSvc{
		repo:         repo,
		productRepo:  productRepo,
		userRepo:     userRepo,
		categoryRepo: categoryRepo,
	}
}

func (svc *transactionSvc) CreateTransaction(ctx context.Context, params *params.CreateTransactions) *views.Response {
	product, err := svc.productRepo.GetProductById(ctx, params.ProductId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, errors.New("product id is invalid"))
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	category, err := svc.categoryRepo.FindCategoryById(ctx, product.CategoryId)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	if product.Stock < params.Quantity {
		return views.ErrorResponse(http.StatusBadRequest, views.M_STOCK_IS_NOT_ENOUGH, errors.New("stock is not enough"))
	}

	userData := ctx.Value("userData").(*common.CustomClaims)

	user, err := svc.userRepo.FindUserById(ctx, uint(userData.Id))
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	if user.Balance < uint(params.Quantity)*product.Price {
		return views.ErrorResponse(http.StatusBadRequest, views.M_BALANCE_IS_NOT_ENOUGH, errors.New("balance is not enough"))
	}

	user.Balance -= uint(params.Quantity) * product.Price
	err = svc.userRepo.UpdateUser(ctx, user)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	product.Stock -= params.Quantity
	err = svc.productRepo.UpdateProduct(ctx, product, product.Id)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	category.SoldProductAmount += uint(params.Quantity)
	err = svc.categoryRepo.UpdateCategory(ctx, category, category.Id)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	model := models.TransactionHistory{
		ProductId:  params.ProductId,
		UserId:     user.Id,
		Quantity:   params.Quantity,
		TotalPrice: params.Quantity * int(product.Price),
	}

	err = svc.repo.CreateTransaction(ctx, &model)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, views.CreateTransaction{
		TotalPrice:   product.Price * uint(params.Quantity),
		Quantity:     uint(params.Quantity),
		ProductTitle: product.Title,
	})
}
