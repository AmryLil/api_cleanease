package orders

import (
	"api_cleanease/features/orders/dtos"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(page, size int) ([]Orders, int64, error)
	Insert(newOrders Orders) error
	SelectByID(ordersID uint) (*Orders, error)
	Update(orders Orders) error
	DeleteByID(ordersID uint) error
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResOrders, int64, error)
	FindByID(ordersID uint) (*dtos.ResOrders, error)
	Create(newOrders dtos.InputOrders) error
	Modify(ordersData dtos.InputOrders, ordersID uint) error
	Remove(ordersID uint) error
}

type Handler interface {
	GetOrderss(c *gin.Context)
	OrdersDetails(c *gin.Context)
	CreateOrders(c *gin.Context)
	UpdateOrders(c *gin.Context)
	DeleteOrders(c *gin.Context)
}
