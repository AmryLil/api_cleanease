package routes

import (
	"api_cleanease/config"
	"api_cleanease/features/orders"
	"api_cleanease/helpers"
	middlewares "api_cleanease/middleware"

	"github.com/gin-gonic/gin"
)

func Orders(r *gin.Engine, handler orders.Handler, jwt helpers.JWTInterface, cfg config.ProgramConfig) {
	orderss := r.Group("/orders")

	orderss.Use(middlewares.AuthorizeJWT(jwt, 1, cfg.SECRET))

	orderss.GET("", handler.GetOrderss)
	orderss.POST("", handler.CreateOrders)

	orderss.GET("/:id", handler.OrdersDetails)
	orderss.PUT("/:id", handler.UpdateOrders)
	orderss.DELETE("/:id", handler.DeleteOrders)
}
