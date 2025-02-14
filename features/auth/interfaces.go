package auth

import (
	"api_cleanease/features/auth/dtos"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(page, size int) ([]User, int64, error)
	Insert(newUser User) error
	SelectByID(userID uint) (*User, error)
	Update(user User) error
	DeleteByID(userID uint) error

	// auth
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResUser, int64, error)
	FindByID(userID uint) (*dtos.ResUser, error)
	Create(newUser dtos.InputUser) error
	Modify(userData dtos.InputUser, userID uint) error
	Remove(userID uint) error

	// auth
}

type Handler interface {
	GetUsers(c *gin.Context)
	UserDetails(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}
