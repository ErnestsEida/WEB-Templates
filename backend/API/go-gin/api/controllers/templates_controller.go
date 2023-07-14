package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	ControllerMethods = append(ControllerMethods, 
		ControllerMethod{ Path: "/", Method: "GET", Handler: handleHome()},
	)
}

func handleHome() gin.HandlerFunc {
	fn  := func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"response": "Template controller is working!",
		})
	}

	return gin.HandlerFunc(fn)
}