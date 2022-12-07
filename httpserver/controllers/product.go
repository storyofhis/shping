package controllers

import (
	"net/http"
	"strconv"

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
	if userData.Role != "admin" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := c.svc.CreateProduct(ctx, &req)
	WriteJsonResponse(ctx, response)
}

func (c *ProductsController) GetProducts(ctx *gin.Context) {
	response := c.svc.GetProducts(ctx)
	WriteJsonResponse(ctx, response)
}

func (c *ProductsController) UpdateProduct(ctx *gin.Context) {
	productId, err := strconv.ParseUint(ctx.Param("productId"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid product id",
		})
		return
	}

	var req params.UpdateProduct
	err = ctx.ShouldBindJSON(&req)
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
	if userData.Role != "admin" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := c.svc.UpdateProduct(ctx, &req, uint(productId))
	WriteJsonResponse(ctx, response)
}

func (c *ProductsController) DeleteProduct(ctx *gin.Context) {
	productId, err := strconv.ParseUint(ctx.Param("productId"), 10, 32)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid product id",
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
	if userData.Role != "admin" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	response := c.svc.DeleteProduct(ctx, uint(productId))
	WriteJsonResponse(ctx, response)
}
