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

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "api_cleanease/docs"

	"github.com/gin-gonic/gin"
)

// @title           CLEANEASE API
// @version         2.0
// @description     API Documentation for Cleanease Laundry Management System
// @termsOfService  http://swagger.io/terms/
// @contact.name    API Support
// @contact.url     http://www.swagger.io/support
// @contact.email   support@cleanease.com
// @license.name    MIT
// @license.url     https://opensource.org/licenses/MIT
// @host            localhost:8001
// @BasePath        /
// @schemes         http https
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// Initialize Gin router
	r := gin.Default()

	// Initialize configuration FIRST
	cfg := config.InitConfig()

	// Configure Swagger info programmatically AFTER getting config
	docs.SwaggerInfo.Title = "CLEANEASE API"
	docs.SwaggerInfo.Description = "API Documentation for Cleanease Laundry Management System"
	docs.SwaggerInfo.Version = "2.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.SERVER_PORT)
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Setup middleware
	middlewares.LogMiddlewares(r)

	// Initialize JWT service
	jwtService := helpers.NewJWT(*cfg)

	// Root endpoint
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to CleanEase API! 😍",
			"version": "2.0",
			"swagger": fmt.Sprintf("http://localhost:%s/swagger/index.html", cfg.SERVER_PORT),
		})
	})

	// Setup routes BEFORE Swagger endpoint
	routes.Orders(r, OrdersHandler(), jwtService, *cfg)
	routes.Packages(r, PackagesHandler())
	routes.Services(r, ServicesHandler())
	routes.Users(r, AuthHandler())

	// Initialize database and run seeders
	db := utils.InitDB()
	seeders.SeedAll(db)

	// Initialize AWS session
	// sess, err := utils.NewSession()
	// if err != nil {
	// 	fmt.Println("Failed to create AWS session:", err)
	// }

	// s3Client := s3.New(sess)
	// fmt.Println("S3 session & client initialized")

	// Swagger endpoint - MUST be AFTER routes setup
	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%s/swagger/doc.json", cfg.SERVER_PORT))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Start server with info
	fmt.Printf("🚀 Server starting on port %s\n", cfg.SERVER_PORT)
	fmt.Printf("📚 Swagger docs: http://localhost:%s/swagger/index.html\n", cfg.SERVER_PORT)
	fmt.Printf("🌐 API Base URL: http://localhost:%s\n", cfg.SERVER_PORT)

	r.Run(fmt.Sprintf(":%s", cfg.SERVER_PORT))
}

func ServicesHandler() laundry_services.Handler {
	db := utils.InitDB()
	db.AutoMigrate(&laundry_services.Services{}) // Use pointer
	repo := sr.New(db)
	usecase := su.New(repo)
	return sh.New(usecase)
}

func PackagesHandler() laundry_packages.Handler {
	db := utils.InitDB()
	db.AutoMigrate(&laundry_packages.Packages{})           // Use pointer
	db.AutoMigrate(&laundry_packages.IndividualPackages{}) // Use pointer

	sess, err := utils.NewSession()
	if err != nil {
		fmt.Println("Failed to create AWS session for packages:", err)
	}

	uploader := s3manager.NewUploader(sess)
	if uploader == nil {
		fmt.Println("Uploader is nil!")
	}

	awsConfig := config.LoadAwsConfig()
	repo := pr.New(db)
	usecase := pu.New(repo)
	return ph.New(usecase, uploader, *awsConfig)
}

func AuthHandler() auth.Handler {
	db := utils.InitDB()
	db.AutoMigrate(&auth.User{})        // Use pointer
	db.AutoMigrate(&auth.UserDetails{}) // Use pointer

	cfg := config.InitConfig()
	jwt := helpers.NewJWT(*cfg)
	hash := helpers.NewHash()
	validation := helpers.NewValidationRequest()
	repo := ar.New(db)
	usecase := au.New(repo, hash, validation, jwt)
	return ah.New(usecase)
}

func OrdersHandler() orders.Handler {
	db := utils.InitDB()
	db.AutoMigrate(&orders.Orders{})    // Use pointer
	db.AutoMigrate(&orders.OrderItem{}) // Use pointer
	repo := or.New(db)
	usecase := ou.New(repo)
	return oh.New(usecase)
}
