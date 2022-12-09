package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/storyofhis/toko-belanja/config"
	"github.com/storyofhis/toko-belanja/httpserver"
	"github.com/storyofhis/toko-belanja/httpserver/controllers"
	"github.com/storyofhis/toko-belanja/httpserver/repositories/gorm"
	"github.com/storyofhis/toko-belanja/httpserver/services"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("cannot load your env")
	}
}

func main() {
	db, err := config.ConnectPostgresGORM()
	if err != nil {
		panic(err)

	}
	router := gin.Default()
	config.GenerateJwtSignature()

	//repo
	userRepo := gorm.NewUserRepo(db)
	categoryRepo := gorm.NewCategoryRepo(db)
	productRepo := gorm.NewProductsRepo(db)
	transactionRepo := gorm.NewTransactionsRepo(db)
	
	// user with admin and customer role
	
	userSvc := services.NewUserSvc(userRepo)
	userHandler := controllers.NewUserController(userSvc)

	// category
	
	categorySvc := services.NewCategorySvc(categoryRepo, productRepo)
	categoryHandler := controllers.NewCategoryController(&categorySvc)

	// product
	
	productSvc := services.NewProductSvc(productRepo, categoryRepo)
	productHandler := controllers.NewProductController(productSvc)

	// transaction
	
	transactionSvc := services.NewTransactionSvc(transactionRepo, productRepo, userRepo, categoryRepo)
	transactionHandler := controllers.NewTransactionController(transactionSvc)

	app := httpserver.NewRouter(router, userHandler, categoryHandler, productHandler, transactionHandler)
	PORT := os.Getenv("PORT")
	app.Start(":" + PORT)
}
