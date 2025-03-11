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
	GetAllIndividualPackages(page, size int) ([]IndividualPackages, int64, error)
	InsertIndividualPackages(newIndividualPackages []IndividualPackages) error
	SelectIndividualPackagesByID(IndividualPackagesID uint) (*IndividualPackages, error)
	UpdateIndividualPackages(IndividualPackages IndividualPackages) error
	DeleteIndividualPackagesByID(IndividualPackagesID uint) error
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResPackages, int64, error)
	FindByID(packagesID uint) (*dtos.ResPackages, error)
	Create(newPackages dtos.InputPackages) error
	Modify(packagesData dtos.InputPackages, packagesID uint) error
	Remove(packagesID uint) error

	// individual package
	FindAllIndividualPackages(page, size int) ([]dtos.ResIndividualPackages, int64, error)
	FindIndividualPackagesByID(IndividualPackagesID uint) (*dtos.ResIndividualPackages, error)
	CreateIndividualPackages(newIndividualPackages []dtos.InputIndividualPackages) error
	ModifyIndividualPackages(IndividualPackagesData dtos.InputIndividualPackages, IndividualPackagesID uint) error
	RemoveIndividualPackages(IndividualPackagesID uint) error
}

type Handler interface {
	GetPackagess(c *gin.Context)
	PackagesDetails(c *gin.Context)
	CreatePackages(c *gin.Context)
	UpdatePackages(c *gin.Context)
	DeletePackages(c *gin.Context)

	// individual package
	GetIndividualPackages(c *gin.Context)
	IndividualPackagesDetails(c *gin.Context)
	CreateIndividualPackages(c *gin.Context)
	UpdateIndividualPackages(c *gin.Context)
	DeleteIndividualPackages(c *gin.Context)
}
