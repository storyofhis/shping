package services

import (
	"context"
	"errors"
	"net/http"

	"github.com/storyofhis/toko-belanja/httpserver/controllers/params"
	"github.com/storyofhis/toko-belanja/httpserver/controllers/views"
	"github.com/storyofhis/toko-belanja/httpserver/repositories"
	"github.com/storyofhis/toko-belanja/httpserver/repositories/models"
	"gorm.io/gorm"
)

type productSvc struct {
	repo         repositories.ProductsRepo
	categoryRepo repositories.CategoryRepo
}

func NewProductSvc(repo repositories.ProductsRepo, categoryRepo repositories.CategoryRepo) ProductSvc {
	return &productSvc{
		repo:         repo,
		categoryRepo: categoryRepo,
	}
}

func (svc *productSvc) CreateProduct(ctx context.Context, product *params.CreateProduct) *views.Response {
	if _, err := svc.categoryRepo.FindCategoryById(ctx, product.CategoryId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, errors.New("category id is invalid"))
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	model := models.Products{
		Title:      product.Title,
		Price:      product.Price,
		Stock:      product.Stock,
		CategoryId: product.CategoryId,
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

func (svc *productSvc) GetProducts(ctx context.Context) *views.Response {
	products, err := svc.repo.GetAllProducts(ctx)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	response := make([]views.GetProduct, 0)
	for i := 0; i < len(products); i++ {
		response = append(response, views.GetProduct{
			Id:         products[i].Id,
			Title:      products[i].Title,
			Price:      products[i].Price,
			Stock:      products[i].Stock,
			CategoryId: products[i].CategoryId,
			CreatedAt:  products[i].CreatedAt,
		})
	}

	return views.SuccessResponse(http.StatusOK, views.M_OK, response)
}

func (svc *productSvc) UpdateProduct(ctx context.Context, params *params.UpdateProduct, productId uint) *views.Response {
	model, err := svc.repo.GetProductById(ctx, productId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, errors.New("product with this id is not found"))
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	if _, err := svc.categoryRepo.FindCategoryById(ctx, params.CategoryId); err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, errors.New("category id is invalid"))
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	model.Title = params.Title
	model.Price = params.Price
	model.Stock = params.Stock
	model.CategoryId = params.CategoryId

	err = svc.repo.UpdateProduct(ctx, model, productId)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_OK, views.UpdateProduct{
		Id:         model.Id,
		Title:      model.Title,
		Price:      model.Price,
		Stock:      model.Stock,
		CategoryId: model.CategoryId,
		CreatedAt:  model.CreatedAt,
		UpdatedAt:  model.UpdatedAt,
	})
}

func (svc *productSvc) DeleteProduct(ctx context.Context, productId uint) *views.Response {
	_, err := svc.repo.GetProductById(ctx, productId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, errors.New("product with this id is not found"))
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	err = svc.repo.DeleteProduct(ctx, productId)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_PRODUCT_SUCCESSFULLY_DELETED, nil)
}
