package routers

import (
	"github.com/aliblue2/khodro45/api/handlers"
	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.RouterGroup) {
	testHandler := handlers.NewTestHandler()
	r.POST("/signup", testHandler.Signup)
	r.POST("/upload-files", testHandler.UploadFileHandler)
}
