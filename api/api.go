package api

import (
	"fmt"

	"github.com/aliblue2/khodro45/api/middlewares"
	"github.com/aliblue2/khodro45/api/routers"
	"github.com/aliblue2/khodro45/api/validations"
	"github.com/aliblue2/khodro45/configs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitiateApp() {
	cfg := configs.GetConfig()
	runingPort := fmt.Sprintf(":%s", cfg.Server.Port)
	r := gin.New()

	validator, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		validator.RegisterValidation("mobile", validations.IranPhoneNumChecker)
	}

	r.Use(gin.Logger(), gin.Recovery(), middlewares.Limiter())
	api := r.Group("/api")
	v1 := api.Group("/v1")

	{
		health := v1.Group("/health")
		test := v1.Group("/test")
		routers.TestRouter(test)
		routers.HealthCheck(health)
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.HealthCheck(health)
	}

	r.Run(runingPort)
}
