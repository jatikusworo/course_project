package common

import (
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

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, BaseResponse{
		StatusCode: "00",
		Message:    "success",
		RefCode:    uuid.New().String(),
		Data:       data,
	})
}

func Error(c *gin.Context, code string, message string) {
	c.JSON(400, BaseResponse{
		StatusCode: code,
		Message:    message,
		RefCode:    uuid.New().String(),
	})
}

func InternalServerError(c *gin.Context, code string, message string) {
	c.JSON(http.StatusInternalServerError, BaseResponse{
		StatusCode: code,
		Message:    message,
		RefCode:    uuid.New().String(),
	})
}
