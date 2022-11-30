package config

import (
	"fmt"
	"os"

	"github.com/storyofhis/toko-belanja/httpserver/repositories/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgresGORM() (*gorm.DB, error) {
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	pass := os.Getenv("PGPASSWORD")
	name := os.Getenv("PGDATABASE")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// automigrate code
	db.Debug().AutoMigrate(models.Users{}, models.Categories{}, models.Products{})

	return db, nil
}
