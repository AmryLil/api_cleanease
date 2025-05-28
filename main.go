package main

import (
	"api_cleanease/config"
	"api_cleanease/database/seeders"
	"api_cleanease/features/auth"
	ah "api_cleanease/features/auth/handler"
	ar "api_cleanease/features/auth/repository"
	au "api_cleanease/features/auth/usecase"
	"api_cleanease/features/laundry_packages"
	ph "api_cleanease/features/laundry_packages/handler"
	pr "api_cleanease/features/laundry_packages/repository"
	pu "api_cleanease/features/laundry_packages/usecase"
	"api_cleanease/features/laundry_services"
	sh "api_cleanease/features/laundry_services/handler"
	sr "api_cleanease/features/laundry_services/repository"
	su "api_cleanease/features/laundry_services/usecase"
	"api_cleanease/features/orders"
	"api_cleanease/helpers"
	middlewares "api_cleanease/middleware"

	oh "api_cleanease/features/orders/handler"
	or "api_cleanease/features/orders/repository"
	ou "api_cleanease/features/orders/usecase"
	"api_cleanease/routes"
	"api_cleanease/utils"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "api_cleanease/docs"

	"github.com/gin-gonic/gin"
)

// @title CleanEase API
// @version 2.0
// @description API for CleanEase laundry management system
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @securityDefinitions.apikey  Bearer
// @in               header
// @name             Authorization
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8001
// @BasePath /
// @schemes http https
func main() {
	r := gin.Default()

	middlewares.LogMiddlewares(r)
	cfg := config.InitConfig()
	jwtService := helpers.NewJWT(*cfg)
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hellooo!😍")
	})
	routes.Orders(r, OrdersHandler(), jwtService, *cfg)
	routes.Packages(r, PackagesHandler())
	routes.Services(r, ServicesHandler())
	routes.Users(r, AuthHandler())

	// seeder
	db := utils.InitDB()
	seeders.SeedAll(db)

	sess, err := utils.NewSession()
	if err != nil {
		fmt.Println("Failed to create AWS session:", err)
	}

	s3Client := s3.New(sess)
	fmt.Println("S3 session & client initialized", s3Client)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(fmt.Sprintf(":%s", cfg.SERVER_PORT))
}

func ServicesHandler() laundry_services.Handler {
	db := utils.InitDB()
	db.AutoMigrate(laundry_services.Services{})
	repo := sr.New(db)
	usecase := su.New(repo)
	return sh.New(usecase)

}

func PackagesHandler() laundry_packages.Handler {
	db := utils.InitDB()
	db.AutoMigrate(laundry_packages.Packages{})
	db.AutoMigrate(laundry_packages.IndividualPackages{})

	sess, _ := utils.NewSession()

	uploader := s3manager.NewUploader(sess)
	if uploader == nil {
		fmt.Println("Uploader is nil!")
	}
	config := config.LoadAwsConfig()
	repo := pr.New(db)
	usecase := pu.New(repo)
	return ph.New(usecase, uploader, *config)

}

func AuthHandler() auth.Handler {
	db := utils.InitDB()
	db.AutoMigrate(auth.User{})
	db.AutoMigrate(auth.UserDetails{})

	config := config.InitConfig()

	jwt := helpers.NewJWT(*config)
	hash := helpers.NewHash()
	validation := helpers.NewValidationRequest()
	repo := ar.New(db)
	usecase := au.New(repo, hash, validation, jwt)
	return ah.New(usecase)
}
func OrdersHandler() orders.Handler {
	db := utils.InitDB()
	db.AutoMigrate(orders.Orders{})
	db.AutoMigrate(orders.OrderItem{})
	repo := or.New(db)
	usecase := ou.New(repo)
	return oh.New(usecase)
}
