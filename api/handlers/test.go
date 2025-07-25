package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
}

type PersonWithNumber struct {
	FirstName string `json:"first_name" binding:"required,min=3,max=20"`
	LastName  string `json:"last_name" binding:"required,min=3,max=20"`
	Number    string `json:"number" binding:"required,min=11,max=11,mobile"`
	Password  string `json:"password" binding:"required,password"`
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (t *TestHandler) Signup(context *gin.Context) {
	person := PersonWithNumber{}
	err := context.ShouldBindJSON(&person)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"person": person})

}

func (t *TestHandler) UploadFileHandler(ctx *gin.Context) {
	file, err := ctx.FormFile("file")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctx.SaveUploadedFile(file, "../forms")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "can not save this file"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successfully uploaded.!"})

}
