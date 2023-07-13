package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TemplateAttachments() []Attachment {
	var returnArray []Attachment

	returnArray = append(returnArray, Attachment{
		Method: "GET",
		Handler: PathHandler{
			Path: "/",
			Function:  handleHome(),
		},
	})

	return returnArray
}

func handleHome() gin.HandlerFunc {
	fn  := func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"response": "Template controller is working!",
		})
	}

	return gin.HandlerFunc(fn)
}