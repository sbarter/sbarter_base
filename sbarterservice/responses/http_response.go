package responses

import (
	"encoding/json"

	"github.com/sbarter/sbarter_be_base_examples/sbartererrors"
	"github.com/sbarter/sbarter_be_base_examples/sbarternetwork"
)

// Response structures a standard HTTP response.
type Response struct {
	Success    bool                      `json:"success" binding:"required"`
	Data       json.RawMessage           `json:"data" binding:"required"`
	Errors     []*ErrorResponse          `json:"errors" binding:"required"`
	ErrorStack *sbartererrors.Error      `json:"errorstack,omitempty"`
	CallStack  *sbarternetwork.CallStack `json:"callstack,omitempty"`
}

// ErrorResponse structures a standard HTTP error response.
type ErrorResponse struct {
	Code    sbartererrors.ErrorCode `json:"code"`
	Message string                  `json:"message"`
	Type    string                  `json:"type"`
}
