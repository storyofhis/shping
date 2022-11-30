package services

import (
	"context"
	"net/http"

	"github.com/storyofhis/toko-belanja/httpserver/controllers/params"
	"github.com/storyofhis/toko-belanja/httpserver/controllers/views"
	"github.com/storyofhis/toko-belanja/httpserver/repositories"
	"github.com/storyofhis/toko-belanja/httpserver/repositories/models"
	"gorm.io/gorm"
)

type productSvc struct {
	repo repositories.ProductsRepo
}

func NewProductSvc(repo repositories.ProductsRepo) ProductSvc {
	return &productSvc{
		repo: repo,
	}
}

func (svc *productSvc) CreateProduct(ctx context.Context, product *params.CreateProduct, categoryId uint) *views.Response {
	if _, err := svc.repo.GetProductById(ctx, product.CategoryId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}
	// request
	model := models.Products{
		Title:      product.Title,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryId: categoryId,
	}

	err := svc.repo.CreateProduct(ctx, &model)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	// response
	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, views.CreateProduct{
		Id:         model.Id,
		Title:      model.Title,
		Price:      model.Price,
		Stock:      model.Stock,
		CategoryId: model.CategoryId,
		CreatedAt:  model.CreatedAt,
	})
}
