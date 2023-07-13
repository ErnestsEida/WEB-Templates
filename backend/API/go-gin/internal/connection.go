package internal

import (
	"gin/api/internal/initializers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB = nil

func OpenConnection() {
	dsn := initializers.DataAsDSN(initializers.DB_Data)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic("Couldn't connect to database!")
	}

	DB = db
}
