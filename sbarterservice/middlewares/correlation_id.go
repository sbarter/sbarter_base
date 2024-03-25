package middlewares

import (
	"strconv"
	"time"

	config "github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// CorrelationID will create a CorrelationID if it doesn't exists. This will allow us to keep
// track of requests and responses as they flow through the distributed architecture.
func CorrelationID(logger *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// If the CorrelationID doesn't exist in the header request, we create it.
		if correlationID := ctx.GetHeader("Correlation-ID"); correlationID == "" {
			cID, cIDErr := uuid.NewRandom()
			if cIDErr != nil {
				if logger != nil {
					logger.WithFields(logrus.Fields{
						"Error": cIDErr,
					}).Error("CorrelationID() failed to generate new UUID")
				}
				return
			}
			correlationID = cID.String()

			ctx.Request.Header.Add("Correlation-ID", correlationID)
		}

		if correlationService := ctx.GetHeader("Correlation-Service"); correlationService == "" {
			correlationService = config.GetString("SERVICE_NAME")

			ctx.Request.Header.Add("Correlation-Service", correlationService)
		}

		if correlationTimeStart := ctx.GetHeader("Correlation-TimeStart"); correlationTimeStart == "" {
			correlationTimeStart = strconv.FormatInt(time.Now().UnixNano(), 10)

			ctx.Request.Header.Add("Correlation-TimeStart", correlationTimeStart)
		}

		ctx.Next()
	}
}
