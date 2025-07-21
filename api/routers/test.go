package routers

import (
	"github.com/aliblue2/khodro45/api/handlers"
	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.RouterGroup) {
	testHandler := handlers.NewTestHandler()

	r.GET("/", testHandler.HeaderBinder)
	r.POST("/form", testHandler.FormBinder)
	r.POST("/form/file", testHandler.FileBinder)
	r.POST("/signup", testHandler.Signup)
}
