package httpserver

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/storyofhis/toko-belanja/common"
	"github.com/storyofhis/toko-belanja/httpserver/controllers"
)

type router struct {
	router *gin.Engine

	user        *controllers.UserController
	category    *controllers.CategoryController
	product     *controllers.ProductsController
	transaction *controllers.TransactionController
}

func NewRouter(r *gin.Engine, user *controllers.UserController, category *controllers.CategoryController, product *controllers.ProductsController, transaction *controllers.TransactionController) *router {
	return &router{
		router:      r,
		user:        user,
		category:    category,
		product:     product,
		transaction: transaction,
	}
}

func (r *router) Start(port string) {
	// users
	r.router.POST("/v1/users/register", r.user.Register)
	r.router.POST("/v1/users/login", r.user.Login)
	r.router.PATCH("/v1/users/topup", r.verifyToken, r.user.TopUp)

	// category
	r.router.POST("/v1/categories", r.verifyToken, r.category.CreateCategory)

	// product
	r.router.POST("/v1/products", r.verifyToken, r.product.CreateProduct)
	r.router.GET("/v1/products", r.verifyToken, r.product.GetProducts)
	r.router.PUT("/v1/products/:productId", r.verifyToken, r.product.UpdateProduct)
	r.router.DELETE("/v1/products/:productId", r.verifyToken, r.product.DeleteProduct)

	// transaction
	r.router.POST("/v1/transactions", r.verifyToken, r.transaction.CreateTransaction)
	r.router.GET("/v1/transactions/my-transactions", r.verifyToken, r.transaction.GetMyTransaction)
	r.router.GET("/v1/transactions/user-transactions", r.verifyToken, r.transaction.GetUserTransaction)

	r.router.Run(port)
}

func (r *router) verifyToken(ctx *gin.Context) {
	bearerToken := strings.Split(ctx.Request.Header.Get("Authorization"), "Bearer ")
	if len(bearerToken) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "invalid bearer token",
		})
		return
	}
	claims, err := common.ValidateToken(bearerToken[1])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.Set("userData", claims)
}
