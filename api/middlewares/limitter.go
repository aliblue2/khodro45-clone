package middlewares

import (
	"github.com/didip/tollbooth/v8"
	"github.com/gin-gonic/gin"
)

func Limiter() gin.HandlerFunc {
	return func(context *gin.Context) {
		limiter := tollbooth.NewLimiter(1, nil)

		httpErr := tollbooth.LimitByRequest(limiter, context.Writer, context.Request)
		if httpErr != nil {
			context.AbortWithStatusJSON(httpErr.StatusCode, gin.H{
				"error": httpErr.Message,
			})
			return
		}
		context.Next()
	}
}
