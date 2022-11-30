package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/storyofhis/toko-belanja/common"
	"github.com/storyofhis/toko-belanja/httpserver/controllers/params"
	"github.com/storyofhis/toko-belanja/httpserver/services"
)

type CategoryController struct {
	svc services.CategorySvc
}

func NewCategoryController(svc *services.CategorySvc) *CategoryController {
	return &CategoryController{
		svc: *svc,
	}
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	// request
	var req params.CreateCategory

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	claims, exists := ctx.Get("userData")
	if !exists {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "token doesn't exists",
		})
		return
	}
	userData := claims.(*common.CustomClaims)
	userRole := userData.Role
	if userRole != "admin" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := c.svc.CreateCategory(ctx, &req)
	WriteJsonResponse(ctx, response)
}
