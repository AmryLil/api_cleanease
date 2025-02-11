package main

import (
	"api_cleanease/config"
	"api_cleanease/features/packages"
	ph "api_cleanease/features/packages/handler"
	pr "api_cleanease/features/packages/repository"
	pu "api_cleanease/features/packages/usecase"
	"api_cleanease/features/services"
	sh "api_cleanease/features/services/handler"
	sr "api_cleanease/features/services/repository"
	su "api_cleanease/features/services/usecase"
	"api_cleanease/routes"
	"api_cleanease/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	cfg := config.InitConfig()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello!😍")
	})
	routes.Packages(r, PackagesHandler())
	routes.Services(r, ServicesHandler())

	r.Run(fmt.Sprintf(":%s", cfg.SERVER_PORT))
}

func ServicesHandler() services.Handler {
	db := utils.InitDB()
	db.AutoMigrate(services.Services{})
	repo := sr.New(db)
	usecase := su.New(repo)
	return sh.New(usecase)

}

func PackagesHandler() packages.Handler {
	db := utils.InitDB()
	db.AutoMigrate(packages.Packages{})
	repo := pr.New(db)
	usecase := pu.New(repo)
	return ph.New(usecase)

}
