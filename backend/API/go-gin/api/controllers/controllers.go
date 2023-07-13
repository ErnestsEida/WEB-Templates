package controllers

import (
	"github.com/gin-gonic/gin"
)

// the method you add to attach controllers to paths, etc. Should return a map that follows this syntax:
// Syntax for attachment passing, you can see the example in --> templates_controller.go
// Example: ControllerAttachment{ Method: "GET", Path: "/", Handler: myFunction() }

var engine *gin.Engine

func AttachControllers(r *gin.Engine) {
	engine = r 		// DO NOT REMOVE
	// ============================

	// ATTACHMENT BLOCK

	attach(TemplateAttachments())
	
	// END OF BLOCK
}

// Auto attachment functionality below

type ControllerAttachment struct {
	Method string
	Path string
	Handler gin.HandlerFunc
}

func attach(methods []ControllerAttachment) {
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