package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Health struct {
}

func NewHealthHandler() *Health {
	return &Health{}
}

func (h *Health) HealthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "its working ..."})
	return
}

func (h *Health) HealthWithId(context *gin.Context) {
	id := context.Params.ByName("id")
	message := fmt.Sprintf("health check handled by id: %s", id)
	context.JSON(http.StatusOK, gin.H{"message": message})
	return
}
