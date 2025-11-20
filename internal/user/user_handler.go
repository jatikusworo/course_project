package user

import (
	"course_project/internal/common"
	"course_project/internal/user/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc Service
}

func NewHandler(s Service) *Handler {
	return &Handler{svc: s}
}

func (h Handler) GetUser(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		common.HandleApiError(context, err)
		return
	}

	u, err := h.svc.GetUser(uint(id))
	if err != nil {
		common.HandleApiError(context, err)
		return
	}

	common.HandleApiSuccess(context, u)
}

func (h Handler) CreateUser(context *gin.Context) {
	var request dto.CreateUserReq
	if err := context.ShouldBindJSON(&request); err != nil {
		common.HandleApiError(context, err)
		return
	}

	u, err := h.svc.CreateUser(request.Name, request.Email)
	if err != nil {
		common.HandleApiError(context, err)
	}

	common.HandleApiSuccess(context, u)
}
