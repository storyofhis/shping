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

type categorySvc struct {
	repo repositories.CategoryRepo
	product repositories.ProductsRepo
}

func NewCategorySvc(repo repositories.CategoryRepo, product repositories.ProductsRepo) CategorySvc {
	return &categorySvc{
		repo: repo,
		product: product,
	}
}

func (svc *categorySvc) CreateCategory(ctx context.Context, category *params.CreateCategory) *views.Response {
	// type
	cat := models.Categories{
		Type: category.Type,
	}

	err := svc.repo.CreateCategory(ctx, &cat)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusCreated, views.M_CREATED, views.CreateCategory{
		Id:                cat.Id,
		Type:              cat.Type,
		SoldProductAmount: cat.SoldProductAmount,
		CreatedAt:         cat.CreatedAt,
	})
}

func (svc *categorySvc) GetCategories(ctx context.Context) *views.Response {
	c, err := svc.repo.GetCategories(ctx)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}
	

	categories := make([]views.GetCategories, 0)
	for _, cat := range c {
		p, err := svc.product.FindProductsByCategoryId(ctx, cat.Id)
			if err != nil {
				return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
				}
		products := make([]views.ProductsGetCategories, 0)
		for _, product := range p {
			products = append(products, views.ProductsGetCategories{
				Id: 	   product.Id,
				Title: 	   product.Title,
				Price: 	   product.Price,
				Stock: 	   product.Stock,
				CreatedAt: product.CreatedAt,
				UpdatedAt: product.UpdatedAt,
			})
		}

		categories = append(categories, views.GetCategories{
			Id: 			   cat.Id,
			Type:		       cat.Type,
			SoldProductAmount: cat.SoldProductAmount,
			CreatedAt: 		   cat.CreatedAt,
			UpdatedAt: 		   cat.UpdatedAt,
			Products:  		   products,
		})
	}
	return views.SuccessResponse(http.StatusOK, views.M_OK, categories)
}

func (svc *categorySvc) UpdateCategory(ctx context.Context, category *params.UpdateCategory, id uint) *views.Response {
	c, err := svc.repo.FindCategoryById(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	c.Type = category.Type

	err = svc.repo.UpdateCategory(ctx, c, id)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_OK, views.UpdateCategory{
		Id:        c.Id,
		Type:      c.Type,
		SoldProductAmount: c.SoldProductAmount,
		UpdatedAt: c.UpdatedAt,
	})
}

func (svc *categorySvc) DeleteCategory(ctx context.Context, id uint) *views.Response {
	_, err := svc.repo.FindCategoryById(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return views.ErrorResponse(http.StatusBadRequest, views.M_BAD_REQUEST, err)
		}
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	err = svc.repo.DeleteCategory(ctx, id)
	if err != nil {
		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
	}

	return views.SuccessResponse(http.StatusOK, views.M_CATEGORY_SUCCESSFULLY_DELETED, nil)
}