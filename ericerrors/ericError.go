package ericerrors

import (
	"net/http"
)

type EricError struct {
	Code    int
	Message string
}

// Custom 404 error
func New404Error(msg string) *EricError {
	return &EricError{
		Code:    http.StatusNotFound,
		Message: msg,
	}
}

// Custom 500 server error
func New500Error(msg string) *EricError {
	return &EricError{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}

// Custom 422  error
func New422Error(msg string) *EricError {
	return &EricError{
		Code:    http.StatusUnprocessableEntity,
		Message: msg,
	}
}
