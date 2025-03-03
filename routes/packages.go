package routes

import (
	"api_cleanease/features/packages"

	"github.com/gin-gonic/gin"
)

func Packages(r *gin.Engine, handler packages.Handler) {
	packagess := r.Group("/packages")

	packagess.GET("", handler.GetPackagess)
	packagess.POST("", handler.CreatePackages)

	packagess.GET("/:id", handler.PackagesDetails)
	packagess.PUT("/:id", handler.UpdatePackages)
	packagess.DELETE("/:id", handler.DeletePackages)

	packagess.GET("/individual", handler.GetIndividualPackages)
	packagess.POST("/individual", handler.CreateIndividualPackages)

	packagess.GET("/individual/:id", handler.IndividualPackagesDetails)
	packagess.PUT("/individual/:id", handler.UpdateIndividualPackages)
	packagess.DELETE("/individual/:id", handler.DeleteIndividualPackages)
}
