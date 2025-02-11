package routes

import (
	"api_cleanease/features/_blueprint"

	"github.com/gin-gonic/gin"
)

func Placeholders(r *gin.Engine, handler _blueprint.Handler) {
	placeholders := r.Group("/placeholder")

	placeholders.GET("", handler.GetPlaceholders)
	placeholders.POST("", handler.CreatePlaceholder)

	placeholders.GET("/:id", handler.PlaceholderDetails)
	placeholders.PUT("/:id", handler.UpdatePlaceholder)
	placeholders.DELETE("/:id", handler.DeletePlaceholder)
}
