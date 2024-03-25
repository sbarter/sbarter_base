package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/sbarter/sbarter_be_base_examples/sbarterutils"

	"github.com/sbarter/sbarter_be_base_examples/sbarternetwork"

	"github.com/gin-gonic/gin"
	"github.com/sbarter/sbarter_be_base_examples/sbartererrors"
	"github.com/sbarter/sbarter_be_base_examples/sbarterservice/responses"
	"github.com/sirupsen/logrus"
)

// BaseHandler sets up the shared abstraction between all handlers.
type baseHandler struct {
	logger         *logrus.Logger
	showErrorStack bool
	showCallStack  bool
}

type BaseHandler interface {
	SendResponse(ctx *gin.Context, data interface{}, errorStack *sbartererrors.Error, callStack *sbarternetwork.CallStack)
}

func NewBaseHandler(logger *logrus.Logger) BaseHandler {
	return &baseHandler{
		logger:         logger,
		showErrorStack: true,
		showCallStack:  true,
	}
}

// SendResponse will expose a generic method to handle sending responses from our endpoints.
// Note that this is a mock method to show the logging format and the response to be shared by all microservices.
func (handler *baseHandler) SendResponse(ctx *gin.Context, data interface{}, errorStack *sbartererrors.Error, callStack *sbarternetwork.CallStack) {

	// Log the errors in the below format that were passed in by other methods, such as handlers, manager, repositories, and services.

	if errorStack != nil {
		log := handler.logger.WithFields(logrus.Fields{
			"StatusCode":      "status code",
			"ClientIP":        "client ip",
			"ClientUserAgent": "client user agent",
			"Referer":         "referer",
			"Service":         "config.GetString('SERVICE_NAME')",
			"CorrelationID":   "ctx.GetHeader('Correlation-ID')",
			"Error":           "errorStack.JSON()",
		})

		log.Error(fmt.Sprintf("%s %s", ctx.Request.Method, ctx.Request.RequestURI))
	}

	// Always mask and omit any sensitive data when logging
	marshalledData, errMarshal := json.Marshal(data)
	if errMarshal != nil {
		if handler.logger != nil {
			handler.logger.WithFields(logrus.Fields{
				"StatusCode":      "status code",
				"ClientIP":        "client ip",
				"ClientUserAgent": "client user agent",
				"Referer":         "referer",
				"Service":         "config.GetString('SERVICE_NAME')",
				"CorrelationID":   "ctx.GetHeader('Correlation-ID')",
				"Data":            sbarterutils.MaskAndOmitObjectForLog(data),
				"Error":           errMarshal,
			}).Error("SendResponse() failed to marshal data to send as response")
		}
	}

	// Control success variable, true/false
	success := true

	// Convert internal status code to HTTP status code
	statusCode := 0

	// Errors
	var errorsResponse []*responses.ErrorResponse

	response := responses.Response{
		Success:    success,
		Data:       marshalledData,
		Errors:     errorsResponse,
		ErrorStack: errorStack,
		CallStack:  callStack,
	}

	// Write CorrelationID to Header
	ctx.Writer.Header().Set("Correlation-ID", ctx.GetHeader("Correlation-ID"))
	ctx.Writer.Header().Set("Correlation-Service", ctx.GetHeader("Correlation-Service"))
	ctx.Writer.Header().Set("Correlation-TimeStart", ctx.GetHeader("Correlation-TimeStart"))

	var duration float64
	var correlationService string
	if correlationService = ctx.GetHeader("Correlation-Service"); correlationService != "" {
		start, errParseInt := strconv.ParseInt(ctx.GetHeader("Correlation-TimeStart"), 10, 64)
		if errParseInt == nil {
			stop := time.Since(time.Unix(0, start))
			duration = float64(stop.Nanoseconds()/1e4) / 100.0
		}
	}

	if handler.logger != nil {
		jsonBytes, _ := json.Marshal(response)
		// Mask sensitive data
		responseForLog := sbarterutils.MaskAndOmitObjectForLog(string(jsonBytes))
		handler.logger.WithFields(logrus.Fields{
			"StatusCode":           ctx.Writer.Status(),
			"ClientIP":             "ctx.ClientIP()",
			"ClientUserAgent":      "ctx.Request.UserAgent()",
			"Referer":              "ctx.Request.Referer()",
			"Service":              "config.GetString('SERVICE_NAME')",
			"CorrelationID":        "ctx.GetHeader('Correlation-ID')",
			"CorrelationService":   "correlationService",
			"CorrelationTimeStart": "ctx.GetHeader('Correlation-TimeStart')",
			"Duration":             duration,
			"Response":             responseForLog,
			"Type":                 "Response",
		}).Debug(fmt.Sprintf("%s %s", ctx.Request.Method, ctx.Request.RequestURI))
	}
	// return response to caller
	ctx.JSON(statusCode, response)
}
