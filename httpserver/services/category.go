package services

import (
	"context"
	"net/http"

	"github.com/storyofhis/toko-belanja/httpserver/controllers/params"
	"github.com/storyofhis/toko-belanja/httpserver/controllers/views"
	"github.com/storyofhis/toko-belanja/httpserver/repositories"
	"github.com/storyofhis/toko-belanja/httpserver/repositories/models"
)

type categorySvc struct {
	repo repositories.CategoryRepo
}

func NewCategorySvc(repo *repositories.CategoryRepo) CategorySvc {
	return &categorySvc{
		repo: *repo,
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

// func (svc *categorySvc) GetCategory(ctx context.Context) *views.Response {
// 	categories, err := svc.repo.GetCategory(ctx)
// 	if err != nil {
// 		return views.ErrorResponse(http.StatusInternalServerError, views.M_INTERNAL_SERVER_ERROR, err)
// 	}

// 	resp := make([]views.GetCategory, len(categories))
// 	for _, cate := range categories {

// 	}
// }
