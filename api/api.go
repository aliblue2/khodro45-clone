package api

import (
	"fmt"

	"github.com/aliblue2/khodro45/api/middlewares"
	"github.com/aliblue2/khodro45/api/routers"
	"github.com/aliblue2/khodro45/api/validations"
	"github.com/aliblue2/khodro45/configs"
	"github.com/aliblue2/khodro45/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitiateServer(cfg *configs.Config) {
	runingPort := fmt.Sprintf(":%s", cfg.Server.InternalPort)
	r := gin.New()

	registerValidators()

	r.Use(middlewares.CorsHandler(*cfg))
	r.Use(middlewares.Limiter())
	r.Use(gin.Logger(), gin.Recovery())

	registerRoutes(r)
	registerSwagger(r, cfg)
	r.Run(runingPort)
}

func registerRoutes(r *gin.Engine) {
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

}

func registerValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		val.RegisterValidation("mobile", validations.IranPhoneNumChecker, true)
		val.RegisterValidation("password", validations.CheckPassword, true)
	}
}

func registerSwagger(r *gin.Engine, cfg *configs.Config) {
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.InternalPort)
	docs.SwaggerInfo.Title = "khodro45"
	docs.SwaggerInfo.Description = "web api for khodro45"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.Version = "1.1"

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
