package error

import "net/http"

type (
	// HandlerError represents an error raised inside a HTTP handler
	HandlerError struct {
		StatusCode int
		Message    string
		Err        error
	}
)

func NewError(statusCode int, message string, err error) *HandlerError {
	return &HandlerError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}

func BadRequest(message string, err error) *HandlerError {
	return NewError(http.StatusBadRequest, message, err)
}

func NotFound(message string, err error) *HandlerError {
	return NewError(http.StatusNotFound, message, err)
}

func InternalServerError(message string, err error) *HandlerError {
	return NewError(http.StatusInternalServerError, message, err)
}

func Unauthorized(message string, err error) *HandlerError {
	return NewError(http.StatusUnauthorized, message, err)
}

func Forbidden(message string, err error) *HandlerError {
	return NewError(http.StatusForbidden, message, err)
}
