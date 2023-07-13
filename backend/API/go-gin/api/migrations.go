package api

import (
	"gin/api/api/models"
	"gin/api/internal/initializers"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate() {
	dsn := initializers.DataAsDSN(initializers.DB_Data)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic("Couldn't connect to database!")
	}

	// Add your migrations here
	db.AutoMigrate(&models.Template{})
}