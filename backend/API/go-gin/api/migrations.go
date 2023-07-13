package api

import (
	"gin/api/api/models"
	"gin/api/internal"
)

func Migrate() {
	// Add your migrations here
	internal.DB.AutoMigrate(&models.Template{})
}