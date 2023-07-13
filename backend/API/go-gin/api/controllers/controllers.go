package controllers

import (
	"github.com/gin-gonic/gin"
)

// the method you add to attach controllers to paths, etc. Should return a map that follows this syntax:
// Syntax for attachment passing, you can see the example in --> templates_controller.go
// { method: 'GET | POST | DELETE | ...', function: { path: "/home | ...", handler: func(ctx *gin.Context) { ... } } }

var engine *gin.Engine

func AttachControllers(r *gin.Engine) {
	engine = r 		// DO NOT REMOVE
	// ============================

	// ATTACHMENT BLOCK

	attach(TemplateAttachments())
	
	// END OF BLOCK
}

// Auto attachment functionality below

type PathHandler struct {
	Path string
	Function gin.HandlerFunc
}

type Attachment struct {
	Method string
	Handler PathHandler
}

func attach(methods []Attachment) {
	for i := 0; i < len(methods); i++ {
		if (methods[i].Method == "GET") {
			engine.GET(methods[i].Handler.Path, methods[i].Handler.Function)
		} else if (methods[i].Method == "POST") {
			engine.POST(methods[i].Handler.Path, methods[i].Handler.Function)
		} else if (methods[i].Method == "DELETE") {
			engine.DELETE(methods[i].Handler.Path, methods[i].Handler.Function)
		} else if (methods[i].Method == "PUT") {
			engine.PUT(methods[i].Handler.Path, methods[i].Handler.Function)
		} else if (methods[i].Method == "PATCH") {
			engine.PATCH(methods[i].Handler.Path, methods[i].Handler.Function)
		} else {
			panic("Incorrect method name was passed and cannot be attached!")
		}
	} 
}