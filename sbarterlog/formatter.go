package sbarterlog

import (
	"time"

	"github.com/sirupsen/logrus"
)

// formatter is used to create a way to add default fields to the logger.
type formatter struct {
	fields logrus.Fields
	lf     logrus.Formatter
}

// Format is used to populate the custom formatter
func (f *formatter) Format(e *logrus.Entry) ([]byte, error) {
	for k, v := range f.fields {
		e.Data[k] = v
	}

	e.Data["Timestamp"] = time.Now().UnixNano()

	return f.lf.Format(e)
}
