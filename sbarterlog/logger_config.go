package sbarterlog

// LoggerConfig is used to pass specific services parameters to the logger
type LoggerConfig struct {
	ServiceName string
	Env         string
	LogLevel    string
}
