package main

import (
	"gin/api/api"
	"gin/api/api/controllers"
	"gin/api/internal"
	"gin/api/internal/initializers"

	"github.com/gin-gonic/gin"
)

// Everything called from "initializers" package should be located in --> /internal/initializers/*.go

func main() {
  internal.OpenConnection()

  if (internal.DB == nil) {
    panic("Database connection was NOT successfully set!")
  } else {  
    api.Migrate() // Migrations are run from --> /api/migrations.go

    r := gin.Default()
    r.Use(initializers.Cors)

    controllers.AttachControllers(r) // Controller attachments are defined in --> /api/controllers/controllers.go
    
    r.Run(initializers.Hostname)
  }
}