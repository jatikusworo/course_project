package common

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)
import "github.com/google/uuid"

type BaseResponse struct {
	StatusCode string      `json:"status_code"`
	Message    string      `json:"message"`
	RefCode    string      `json:"ref_code,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func HandleApiSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, BaseResponse{
		StatusCode: "00",
		Message:    "success",
		RefCode:    uuid.New().String(),
		Data:       data,
	})
}

func HandleApiError(c *gin.Context, err error) {
	var apiErr *ApiError
	if errors.As(err, &apiErr) {
		Error(c, apiErr.HttpStatus, apiErr.StatusCode, apiErr.Message)
		return
	}

	// fallback: unexpected error
	Error(c, http.StatusInternalServerError, "99", err.Error())
}

func Error(c *gin.Context, httpStatus int, code string, message string) {
	c.JSON(400, BaseResponse{
		StatusCode: code,
		Message:    message,
		RefCode:    uuid.New().String(),
	})
}
