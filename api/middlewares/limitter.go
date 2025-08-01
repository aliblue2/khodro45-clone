package middlewares

import (
	"net/http"

	"github.com/aliblue2/khodro45/api/helper"
	"github.com/didip/tollbooth/v8"
	"github.com/gin-gonic/gin"
)

func Limiter() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)
	return func(ctx *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, helper.BaseResponseHandlerWithError(nil, false, 0, err))
			return
		} else {
			ctx.Next()
		}
	}

}
