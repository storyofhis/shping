package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/storyofhis/toko-belanja/common"
	"github.com/storyofhis/toko-belanja/httpserver/controllers/params"
	"github.com/storyofhis/toko-belanja/httpserver/services"
)

type ProductsController struct {
	svc services.ProductSvc
}

func NewProductController(svc services.ProductSvc) *ProductsController {
	return &ProductsController{
		svc: svc,
	}
}

func (c *ProductsController) CreateProduct(ctx *gin.Context) {
	var req params.CreateProduct

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
	UserID := userData.Id
	// userRole := userData.Role
	// if userRole != "admin" {
	// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
	// 		"error": "unauthorized",
	// 	})
	// 	return
	// }
	err = validator.New().Struct(req)

	response := c.svc.CreateProduct(ctx, &req, uint(UserID))
	WriteJsonResponse(ctx, response)
}
