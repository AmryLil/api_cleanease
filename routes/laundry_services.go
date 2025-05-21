package routes

import (
	"api_cleanease/features/laundry_services"

	"github.com/gin-gonic/gin"
)

func Services(r *gin.Engine, handler laundry_services.Handler) {
	servicess := r.Group("/services")

	servicess.GET("", handler.GetServicess)
	servicess.POST("", handler.CreateServices)

	servicess.GET("/:id", handler.ServicesDetails)
	servicess.PUT("/:id", handler.UpdateServices)
	servicess.DELETE("/:id", handler.DeleteServices)
}
