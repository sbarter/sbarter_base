package middlewares

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	config "github.com/spf13/viper"
)

// RequestLogger will log every incoming request
func RequestLogger(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if logger != nil && ctx.Request != nil {
			// Ignore OPTIONS requests
			if ctx.Request.Method == "OPTIONS" {
				return
			}

			// Read the request body and create 2 buffers from it: one for the logs, one for the actual request
			buf, _ := io.ReadAll(ctx.Request.Body)
			logBuffer := io.NopCloser(bytes.NewBuffer(buf))
			buffer := io.NopCloser(bytes.NewBuffer(buf))

			logger.WithFields(logrus.Fields{
				"Service":       config.GetString("SERVICE_NAME"),
				"CorrelationID": ctx.GetHeader("Correlation-ID"),
				"Payload":       readBodyAndMask(logBuffer),
				"Type":          "Request",
				"Method":        ctx.Request.Method,
				"URI":           ctx.Request.RequestURI,
			}).Debug(fmt.Sprintf("%s %s", ctx.Request.Method, ctx.Request.RequestURI))

			ctx.Request.Body = buffer
		}

		ctx.Next()
	}
}

func readBodyAndMask(reader io.Reader) string {
	// Create new buffer, mask bytes and return
	return ""
}
