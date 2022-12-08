package controllers

import (
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
	response := c.svc.GetUserTransaction(ctx)
	WriteJsonResponse(ctx, response)
}
