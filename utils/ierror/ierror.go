package ierror

import (
	"fmt"
)

// IError defines the company specific error interface
type IError interface {
	error
	GetCode() uint64
	AddMsg(args ...interface{}) IError
}

// HTTPError defines the  error contract for HTTP requests.
// It implements the IError interface
type HTTPError struct {
	HTTPCode    uint64 `json:"http_code"`
	Code        uint64 `json:"code"`
	Description string `json:"description"`
	Message     string `json:"message"`
}

// New gives back a new instance of HTTPError
func New(httpCode, errCode uint64, description, message string) func(...interface{}) IError {
	return func(args ...interface{}) IError {
		return newHTTPErr(httpCode, errCode, description, message).AddMsg(args...)
	}
}

func newHTTPErr(httpCode, errCode uint64, description, message string) *HTTPError {
	return &HTTPError{
		HTTPCode:    httpCode,
		Code:        errCode,
		Description: description,
		Message:     message,
	}
}

// Error implements the stringify method of error interface
func (e HTTPError) Error() string {
	return e.Message
}

// GetCode returns Error code from HTTPError
func (e HTTPError) GetCode() uint64 { return e.Code }

// AddMsg adds a message to HTTPError.
// This returns the error itself.
func (e *HTTPError) AddMsg(args ...interface{}) IError {
	if args != nil {
		var m string
		for _, v := range args {
			m = fmt.Sprintf("%s %v", m, v)
		}
		e.Message = fmt.Sprintf("%s%v", e.Message, m)
	}
	return e

}
