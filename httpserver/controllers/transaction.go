package controllers

import (
	"github.com/storyofhis/toko-belanja/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/storyofhis/toko-belanja/httpserver/controllers/params"
	"github.com/storyofhis/toko-belanja/httpserver/services"
)

type TransactionController struct {
	svc services.TransactionSvc
}

func NewTransactionController(svc services.TransactionSvc) *TransactionController {
	return &TransactionController{
		svc: svc,
	}
}

func (c *TransactionController) CreateTransaction(ctx *gin.Context) {
	var req params.CreateTransactions

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
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

	response := c.svc.CreateTransaction(ctx, &req)
	WriteJsonResponse(ctx, response)
}

func (c *TransactionController) GetMyTransaction(ctx *gin.Context) {
	response := c.svc.GetMyTransaction(ctx)
	WriteJsonResponse(ctx, response)
}

func (c *TransactionController) GetUserTransaction(ctx *gin.Context) {
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
			"error": "unauthorized user",
		})
		return
	}

	response := c.svc.GetUserTransaction(ctx)
	WriteJsonResponse(ctx, response)
}
