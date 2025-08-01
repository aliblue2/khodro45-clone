package handlers

import (
	"fmt"
	"net/http"

	"github.com/aliblue2/khodro45/api/helper"
	"github.com/gin-gonic/gin"
)

type Health struct {
}

func NewHealthHandler() *Health {
	return &Health{}
}

// @Summary health check of v1 api
// @Description health check by calling a get request
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} helper.BaseResponse "success"
// @Failure 400 {object} helper.BaseResponse "failed"
// @Router /v1/health [get]
func (h *Health) HealthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, helper.BaseResponseHandler("im working", true, 1))
	return
}

// @Summary health check of v1 api
// @Description health check by calling a get request
// @Tags health
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} helper.BaseResponse "success"
// @Failure 400 {object} helper.BaseResponse "failed"
// @Router /v1/health/{id} [get]
func (h *Health) HealthWithId(context *gin.Context) {
	id := context.Params.ByName("id")
	message := fmt.Sprintf("health check handled by id: %s", id)
	context.JSON(http.StatusOK, helper.BaseResponseHandler(message, true, 1))
	return
}
