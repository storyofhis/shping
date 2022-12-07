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

	// user with admin and customer role
	userRepo := gorm.NewUserRepo(db)
	userSvc := services.NewUserSvc(userRepo)
	userHandler := controllers.NewUserController(userSvc)

	// category
	categoryRepo := gorm.NewCategoryRepo(db)
	categorySvc := services.NewCategorySvc(&categoryRepo)
	categoryHandler := controllers.NewCategoryController(&categorySvc)

	// product
	productRepo := gorm.NewProductsRepo(db)
	productSvc := services.NewProductSvc(productRepo)
	productHandler := controllers.NewProductController(productSvc)

	app := httpserver.NewRouter(router, userHandler, categoryHandler, productHandler)
	PORT := os.Getenv("PORT")
	app.Start(":" + PORT)
}
