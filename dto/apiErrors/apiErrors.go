package apiErrors

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (r *CustomError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Message)
}

func BadRequest(message string) *CustomError {
	return &CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}

func Unauthorized(message string) *CustomError {
	return &CustomError{
		StatusCode: http.StatusUnauthorized,
		Message:    message,
	}
}

func NotFound(message string) *CustomError {
	return &CustomError{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

func Conflict(message string) *CustomError {
	return &CustomError{
		StatusCode: http.StatusConflict,
		Message:    message,
	}
}

