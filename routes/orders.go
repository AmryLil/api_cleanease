package routes

import (
	"api_cleanease/features/orders"

	"github.com/gin-gonic/gin"
)

func Orderss(r *gin.Engine, handler orders.Handler) {
	orderss := r.Group("/orders")

	orderss.GET("", handler.GetOrderss)
	orderss.POST("", handler.CreateOrders)

	orderss.GET("/:id", handler.OrdersDetails)
	orderss.PUT("/:id", handler.UpdateOrders)
	orderss.DELETE("/:id", handler.DeleteOrders)
}
