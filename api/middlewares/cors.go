package middlewares

import (
	"github.com/aliblue2/khodro45/configs"
	"github.com/gin-gonic/gin"
)

func CorsHandler(cfg configs.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		allowOrigin := cfg.Cors.AllowOrigins

		if allowOrigin != "" {
			ctx.Header("Access-Control-Allow-Origin", allowOrigin)
			ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
			ctx.Header("Access-Control-Allow-Credentials", "true")
		}

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}
		ctx.Next()
	}
}
