package sbarterutils

const (
	defaultMaskString = "***"
	defaultOmitString = "***omitted***"
)

// Configurable sensitive fields to mask; This means that all properties must be standardized and not having phone in one service and phoneNum in another service.
var defaultSensitiveFields = []string{
	"name",
	"surname",
	"nationality",
	"card",
	"phone",
	"username",
	"password",
	"email",
	"address",
}

// MaskAndOmitObjectForLog, will take care of masking sensitive data and omit properties like files.
func MaskAndOmitObjectForLog(v interface{}) string {
	return ""
}
