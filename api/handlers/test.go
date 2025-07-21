package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestHandler struct {
}

type Header struct {
	userId      int64
	browserName string
}

type Person struct {
	FirstName string
	LastName  string
}

type PersonWithNumber struct {
	FirstName string `json:"first_name" binding:"required,min=3,max=20"`
	LastName  string `json:"last_name" binding:"required,min=3,max=20"`
	Number    string `json:"number" binding:"required,min=11,max=11,mobile"`
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (t *TestHandler) HeaderBinder(context *gin.Context) {
	tmpHeader := Header{}
	err := context.BindHeader(&tmpHeader)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "can not bind keys and values"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"userId": tmpHeader.userId, "browserName": tmpHeader.browserName})

}

func (t *TestHandler) FormBinder(context *gin.Context) {
	tmpPerson := Person{}
	err := context.ShouldBind(&tmpPerson)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "cant not bind data from form data"})
	}

	context.JSON(http.StatusOK, gin.H{"person": tmpPerson})

}

func (t *TestHandler) FileBinder(context *gin.Context) {
	file, err := context.FormFile("file")

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = context.SaveUploadedFile(file, "../forms")

	if err != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"fileName": file.Filename, "size": file.Size})
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
