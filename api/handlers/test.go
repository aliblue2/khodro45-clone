package handlers

import (
	"net/http"

	"github.com/aliblue2/khodro45/api/helper"
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

// Signup
// @Summary signup
// @Description signup user with mobile number
// @Tags signup
// @Accept json
// @Produce json
// @Param person body PersonWithNumber true "person"
// @Success 200 {PersonWithNumber} helper.BaseResponse "success"
// @Failure 400 {object} helper.BaseResponse "failed"
// @Failure 500 {object} helper.BaseResponse "failed"
// @Router /v1/test/signup [post]
func (t *TestHandler) Signup(context *gin.Context) {
	person := PersonWithNumber{}
	err := context.ShouldBindJSON(&person)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, helper.BaseResponseHandlerWithValidationError("failed", false, 0, err))
		return
	}

	context.JSON(http.StatusOK, helper.BaseResponseHandler(person, true, 1))

}

// FileUpload
// @Summary file upload
// @Description upload file to server
// @Tags upload
// @Accept json
// @Produce json
// @Param file formData file true "file"
// @Success 200 {object} helper.BaseResponse "success"
// @Failure 400 {object} helper.BaseResponse "failed"
// @Failure 500 {object} helper.BaseResponse "failed"
// @Router /v1/test/upload-files [post]
func (t *TestHandler) UploadFileHandler(ctx *gin.Context) {
	file, err := ctx.FormFile("file")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.BaseResponseHandlerWithError("failed", false, 0, err))
		return
	}

	err = ctx.SaveUploadedFile(file, "../forms")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.BaseResponseHandlerWithError("can not save this file", false, 0, err))
		return
	}

	ctx.JSON(http.StatusOK, helper.BaseResponseHandler("file successfully uploaded.!", true, 1))

}
