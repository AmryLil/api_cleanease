package packages

import (
	"api_cleanease/features/packages/dtos"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	GetAll(page, size int) ([]Packages, int64, error)
	Insert(newPackages Packages) error
	SelectByID(packagesID uint) (*Packages, error)
	Update(packages Packages) error
	DeleteByID(packagesID uint) error
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResPackages, int64, error)
	FindByID(packagesID uint) (*dtos.ResPackages, error)
	Create(newPackages dtos.InputPackages) error
	Modify(packagesData dtos.InputPackages, packagesID uint) error
	Remove(packagesID uint) error
}

type Handler interface {
	GetPackagess(c *gin.Context)
	PackagesDetails(c *gin.Context)
	CreatePackages(c *gin.Context)
	UpdatePackages(c *gin.Context)
	DeletePackages(c *gin.Context)
}
