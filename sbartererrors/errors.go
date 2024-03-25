package sbartererrors

type ErrorCode int

const (
	// Examples of internal Error Codes to be used throughout the system
	GenericError         ErrorCode = 2000
	ValidationError      ErrorCode = 2001
	DatabaseError        ErrorCode = 2002
	RecordNotFoundError  ErrorCode = 2003
	CacheError           ErrorCode = 2004
	MaintenanceErrorCode ErrorCode = 2005
	FileStorageError     ErrorCode = 2006
	ProviderError        ErrorCode = 2007
)

// Error holds the error information together with all previous errors
type Error struct {
	Code        ErrorCode    `json:"code"`
	Message     string       `json:"message"`
	Timestamp   int64        `json:"timestamp"`
	IsSticky    bool         `json:"sticky,omitempty"`
	RuntimeInfo *RuntimeInfo `json:"runtimeinfo,omitempty"`
	Previous    *Error       `json:"previous,omitempty"`
}

type RuntimeInfo struct {
	Project  string `json:"project"`
	Function string `json:"function"`
	Path     string `json:"path"`
	Line     int    `json:"line"`
}
