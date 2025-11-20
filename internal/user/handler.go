package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc Service
}

func NewHandler(s Service) *Handler {
	return &Handler{svc: s}
}

func RegisterRoutes(r *gin.Engine, h *Handler) {
	g := r.Group("/user")
	g.GET(":id", h.GetUser)
	g.POST("", h.CreateUser)
}

func (h Handler) GetUser(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id " + err.Error()})
		return
	}

	u, err := h.svc.GetUser(uint(id))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if u == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	context.JSON(http.StatusOK, u)
}

func (h Handler) CreateUser(context *gin.Context) {
	var request CreateUserReq
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.svc.CreateUser(request.Name, request.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	if u == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Save User Failed"})
	}

	context.JSON(http.StatusOK, u)
}
