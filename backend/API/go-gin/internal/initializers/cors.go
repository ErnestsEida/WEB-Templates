package initializers

import "github.com/gin-contrib/cors"

var Cors = cors.New(cors.Config{
	AllowOrigins:     []string{"*"},
	AllowMethods:     []string{"*"},
	AllowCredentials: true,
})