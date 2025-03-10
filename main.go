package main

import (
	"api_cleanease/config"
	"api_cleanease/features/auth"
	ah "api_cleanease/features/auth/handler"
	ar "api_cleanease/features/auth/repository"
	au "api_cleanease/features/auth/usecase"
	"api_cleanease/features/orders"
	"api_cleanease/features/packages"
	ph "api_cleanease/features/packages/handler"
	pr "api_cleanease/features/packages/repository"
	pu "api_cleanease/features/packages/usecase"
	"api_cleanease/features/services"
	sh "api_cleanease/features/services/handler"
	sr "api_cleanease/features/services/repository"
	su "api_cleanease/features/services/usecase"
	"api_cleanease/helpers"

	oh "api_cleanease/features/orders/handler"
	or "api_cleanease/features/orders/repository"
	ou "api_cleanease/features/orders/usecase"
	"api_cleanease/routes"
	"api_cleanease/utils"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	cfg := config.InitConfig()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello!😍")
	})
	routes.Orders(r, OrdersHandler())
	routes.Packages(r, PackagesHandler())
	routes.Services(r, ServicesHandler())
	routes.Users(r, AuthHandler())

	sess, err := utils.NewSession()
	if err != nil {
		fmt.Println("Failed to create AWS session:", err)
	}

	s3Client := s3.New(sess)
	fmt.Println("S3 session & client initialized", s3Client)

	r.Run(fmt.Sprintf(":%s", cfg.SERVER_PORT))
}

func ServicesHandler() services.Handler {
	db := utils.InitDB()
	repo := sr.New(db)
	usecase := su.New(repo)
	return sh.New(usecase)

}

func PackagesHandler() packages.Handler {
	db := utils.InitDB()
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
	repo := or.New(db)
	usecase := ou.New(repo)
	return oh.New(usecase)
}
