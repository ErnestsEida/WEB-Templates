package controllers

import (
	"github.com/gin-gonic/gin"
)

// This is a ready file for usage, that handles controller attaching to gin engine
// You attach your controller methods by adding your method that returns "gin.HandlerFunc", using the ControllerMethod struct (see template for example)

type ControllerMethod struct {
	Method string
	Path string
	Handler gin.HandlerFunc
}

var ControllerMethods []ControllerMethod
var engine *gin.Engine

func AttachControllers(r *gin.Engine) {
	engine = r 		// DO NOT REMOVE
	ProcessControllerMethods()
}

// Auto attachment functionality below
func ProcessControllerMethods() {
	methods := ControllerMethods
	for i := 0; i < len(methods); i++ {
		if (methods[i].Method == "GET") {
			engine.GET(methods[i].Path, methods[i].Handler)
		} else if (methods[i].Method == "POST") {
			engine.POST(methods[i].Path, methods[i].Handler)
		} else if (methods[i].Method == "DELETE") {
			engine.DELETE(methods[i].Path, methods[i].Handler)
		} else if (methods[i].Method == "PUT") {
			engine.PUT(methods[i].Path, methods[i].Handler)
		} else if (methods[i].Method == "PATCH") {
			engine.PATCH(methods[i].Path, methods[i].Handler)
		} else {
			panic("Incorrect method name was passed and cannot be attached!")
		}
	} 
}