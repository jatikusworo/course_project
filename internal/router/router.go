package router

import (
	"course_project/internal/user"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	User *user.Handler
	//Room *room.Handler
}

func RegisterRoutes(r *gin.Engine, h *Handlers) {

	userGroup := r.Group("/user")
	{
		userGroup.POST("", h.User.CreateUser)
		userGroup.GET("/:id", h.User.GetUser)
	}

	//roomGroup := r.Group("/room")
	//{
	//	roomGroup.POST("", h.Room.CreateRoom)
	//	roomGroup.GET("/:id", h.Room.GetRoom)
	//	roomGroup.GET("", h.Room.ListRooms)
	//}
}
