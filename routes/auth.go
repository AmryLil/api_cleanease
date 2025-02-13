package routes

import (
	user "api_cleanease/features/auth"

	"github.com/gin-gonic/gin"
)

func Users(r *gin.Engine, handler user.Handler) {
	users := r.Group("/auth")

	users.GET("", handler.GetUsers)
	users.POST("", handler.CreateUser)

	users.GET("/:id", handler.UserDetails)
	users.PUT("/:id", handler.UpdateUser)
	users.DELETE("/:id", handler.DeleteUser)
}
