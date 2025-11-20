package common

import (
	"net/http"

	"github.com/google/uuid"
)

type ApiError struct {
	HttpStatus int    `json:"http_status"`
	StatusCode string `json:"status_code"`
	Message    string `json:"message"`
	RefCode    string `json:"ref_code"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewBadRequest(statusCode, message string) error {
	return &ApiError{
		HttpStatus: http.StatusBadRequest,
		StatusCode: statusCode,
		Message:    message,
		RefCode:    uuid.New().String(),
	}
}

func NewNotFound(message string) error {
	return &ApiError{
		HttpStatus: ErrNotFound.HttpStatus(),
		StatusCode: ErrNotFound.StatusCode(),
		Message:    message,
		RefCode:    uuid.New().String(),
	}
}

func NewForbidden(statusCode, message string) error {
	return &ApiError{
		HttpStatus: http.StatusForbidden,
		StatusCode: statusCode,
		Message:    message,
		RefCode:    uuid.New().String(),
	}
}

func NewInternalError(message string) error {
	return &ApiError{
		HttpStatus: ErrUnknown.HttpStatus(),
		StatusCode: ErrUnknown.StatusCode(),
		Message:    message,
		RefCode:    uuid.New().String(),
	}
}
