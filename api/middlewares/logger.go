package middlewares

import (
	"bytes"
	"io"
	"time"

	"github.com/aliblue2/khodro45/configs"
	"github.com/aliblue2/khodro45/pkg/logging"
	"github.com/gin-gonic/gin"
)

type bodyLoggerWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyLoggerWriter) write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w *bodyLoggerWriter) writeString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func DefaultStructuredLogger(cfg *configs.Config) gin.HandlerFunc {
	logger := logging.NewLogger(cfg)
	return structuredLogger(logger)
}

func structuredLogger(logger logging.Logger) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		blw := &bodyLoggerWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		start := time.Now() /// start time
		path := ctx.FullPath()
		rawQuery := ctx.Request.URL.RawQuery

		bodyBytes, err := io.ReadAll(ctx.Request.Body)

		if err != nil {
			logger.Error(logging.RequestResponse, logging.Api, "failed to read context body", nil)
		}

		ctx.Request.Body.Close()
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		ctx.Writer = blw
		ctx.Next()

		param := gin.LogFormatterParams{}
		param.TimeStamp = time.Now() /// end request time
		param.Latency = param.TimeStamp.Sub(start)
		param.ClientIP = ctx.ClientIP()
		param.Method = ctx.Request.Method
		param.StatusCode = ctx.Writer.Status()
		param.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = ctx.Writer.Size()

		if rawQuery != "" {
			path = path + "?" + rawQuery
		}

		param.Path = path

		keys := map[logging.ExtraKey]interface{}{}
		keys[logging.ClientIp] = param.ClientIP
		keys[logging.Method] = param.Method
		keys[logging.Latency] = param.Latency
		keys[logging.StatusCode] = param.StatusCode
		keys[logging.ErrorMessage] = param.ErrorMessage
		keys[logging.BodySize] = param.BodySize
		keys[logging.RequestBody] = string(bodyBytes)
		keys[logging.ResponseBody] = blw.body.String()

		logger.Info(logging.RequestResponse, logging.Api, "", keys)
	}

}
