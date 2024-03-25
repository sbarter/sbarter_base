package sbarterlog

import (
	"github.com/sirupsen/logrus"
)

// NewLogger is a constructor that returns a reference to a new Logger instance.
// This will always log the Service and the environment ( e.g., local, development, staging & production )
func NewLogger(config LoggerConfig) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&formatter{
		fields: logrus.Fields{
			"Service": config.ServiceName,
			"Env":     config.Env,
		},
		lf: &logrus.JSONFormatter{},
	})

	if logLevel, err := logrus.ParseLevel(config.LogLevel); err == nil {
		logger.SetLevel(logLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	return logger
}
