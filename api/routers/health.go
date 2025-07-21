package routers

import (
	"github.com/aliblue2/khodro45/api/handlers"
	"github.com/gin-gonic/gin"
)

func HealthCheck(r *gin.RouterGroup) {
	health := handlers.NewHealthHandler()
	r.GET("/", health.HealthCheck)
	r.GET("/:id", health.HealthWithId)
}
