package _blueprint

import (
	"api_cleanease/features/_blueprint/dtos"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(page, size int) ([]Placeholder, int64, error)
	Insert(newPlaceholder Placeholder) error
	SelectByID(placeholderID uint) (*Placeholder, error)
	Update(placeholder Placeholder) error
	DeleteByID(placeholderID uint) error
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResPlaceholder, int64, error)
	FindByID(placeholderID uint) (*dtos.ResPlaceholder, error)
	Create(newPlaceholder dtos.InputPlaceholder) error
	Modify(placeholderData dtos.InputPlaceholder, placeholderID uint) error
	Remove(placeholderID uint) error
}

type Handler interface {
	GetPlaceholders(c *gin.Context)
	PlaceholderDetails(c *gin.Context)
	CreatePlaceholder(c *gin.Context)
	UpdatePlaceholder(c *gin.Context)
	DeletePlaceholder(c *gin.Context)
}
