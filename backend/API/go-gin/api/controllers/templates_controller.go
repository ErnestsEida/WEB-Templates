package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TemplateAttachments() []ControllerAttachment {
	return []ControllerAttachment{
		{ Method: "GET", Path: "/", Handler: handleHome() },
	}
}

func handleHome() gin.HandlerFunc {
	fn  := func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"response": "Template controller is working!",
		})
	}

	return gin.HandlerFunc(fn)
}