package api

import (
	"gin/api/api/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate() {
	dsn := "user=postgres password=postgres dbname=go_gin_api_db host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic("Couldn't connect to database!")
	}

	// Add your migrations here
	db.AutoMigrate(&models.Template{})
}