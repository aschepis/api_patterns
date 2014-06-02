package error_handling

import (
	"fmt"
	"net/http"

	"github.com/aschepis/api_patterns/rendering"
)

// error type suitable for returning to
type APIError struct {
	err        error
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

// a function type that returns an APIError
type APIErrorFunc func() *APIError

// wrapper to generate api error for internal errors
func InternalError() *APIError {
	return &APIError{
		Message:    fmt.Sprintf("Internal Server Error"),
		StatusCode: http.StatusInternalServerError,
	}
}

// wrapper to generate api error for Forbidden errors
func ForbiddenError() *APIError {
	return &APIError{
		Message:    fmt.Sprintf("Forbidden"),
		StatusCode: http.StatusForbidden,
	}
}

// helper for generating api error objects
func MakeAPIError(err error, code int) *APIError {
	return &APIError{
		Message:    err.Error(),
		StatusCode: code,
	}
}

// Handy wrapper function to help with giving decent error responses/codes
// when an error occurs.
func WrapError(w http.ResponseWriter, r rendering.RenderFunc, f APIErrorFunc) {
	if err := f(); err != nil {
		w.WriteHeader(err.StatusCode)
		r(w, err)
	}
}
