package services

import (
	"api_cleanease/features/services/dtos"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(page, size int) ([]Services, int64, error)
	Insert(newServices Services) error
	SelectByID(servicesID uint) (*Services, error)
	Update(services Services) error
	DeleteByID(servicesID uint) error
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResServices, int64, error)
	FindByID(servicesID uint) (*dtos.ResServices, error)
	Create(newServices dtos.InputServices) error
	Modify(servicesData dtos.InputServices, servicesID uint) error
	Remove(servicesID uint) error
}

type Handler interface {
	GetServicess(c *gin.Context)
	ServicesDetails(c *gin.Context)
	CreateServices(c *gin.Context)
	UpdateServices(c *gin.Context)
	DeleteServices(c *gin.Context)
}
