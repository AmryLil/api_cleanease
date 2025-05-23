package handler

import (
	"api_cleanease/config"
	"api_cleanease/features/laundry_packages"
	"api_cleanease/features/laundry_packages/dtos"
	"api_cleanease/helpers"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	service  laundry_packages.Usecase
	uploader *s3manager.Uploader
	config   config.AWSConfig
}

func New(service laundry_packages.Usecase, uploader *s3manager.Uploader, config config.AWSConfig) laundry_packages.Handler {
	return &controller{
		service:  service,
		uploader: uploader,
		config:   config,
	}
}

var validate *validator.Validate

// GetLaundryPackages godoc
// @Summary Get all laundry packages
// @Description Get all laundry packages with pagination
// @Tags Laundry Packages
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1) minimum(1)
// @Param size query int false "Page size" default(5) minimum(1)
// @Success 200 {object} helpers.ResponseGetAllSuccess{data=[]dtos.ResPackages,pagination=helpers.Pagination} "Get all packages success"
// @Failure 400 {object} helpers.ResponseError "Invalid pagination data"
// @Failure 404 {object} helpers.ResponseError "No packages found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /packages [get]
func (ctl *controller) CreatePackages(c *gin.Context) {
	file, err := c.FormFile("cover")
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("File cover is required"))
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse("Failed to open file"))
		return
	}
	defer src.Close()

	err = helpers.UploadFileFromReader(ctl.uploader, src, ctl.config.S3Bucket, file.Filename, file.Size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse("Failed to upload to S3 "+err.Error()))
		return
	}

	imageURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", ctl.config.S3Bucket, ctl.config.Region, file.Filename)

	name := c.PostForm("name")
	pricePerKgStr := c.PostForm("price_per_kg")
	description := c.PostForm("description")

	pricePerKg, err := strconv.ParseFloat(pricePerKgStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Invalid price_per_kg format"))
		return
	}

	input := dtos.InputPackages{
		Name:        name,
		PricePerKg:  pricePerKg,
		Description: description,
		Cover:       imageURL,
	}

	err = ctl.service.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: "Create Package Success",
		Status:  true,
	})
}

func (ctl *controller) GetPackagess(c *gin.Context) {
	var pagination dtos.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Please provide valid pagination data!"))
		return
	}

	if pagination.Page < 1 || pagination.Size < 1 {
		pagination.Page = 1
		pagination.Size = 5
	}
	page := pagination.Page
	pageSize := pagination.Size

	packagess, total, err := ctl.service.FindAll(page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if packagess == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("There is No Packagess!"))
		return
	}

	paginationData := helpers.PaginationResponse(page, pageSize, int(total))

	c.JSON(http.StatusOK, helpers.ResponseGetAllSuccess{
		Status:     true,
		Message:    "Get All Packagess Success",
		Data:       packagess,
		Pagination: paginationData,
	})
}

func (ctl *controller) PackagesDetails(c *gin.Context) {
	packagesID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	packages, err := ctl.service.FindByID(uint(packagesID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if packages == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Packages Not Found!"))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseGetDetailSuccess{
		Data:    packages,
		Status:  true,
		Message: " Get Packages Detail Success",
	})
}

func (ctl *controller) UpdatePackages(c *gin.Context) {
	var input dtos.InputPackages
	packagesID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	packages, err := ctl.service.FindByID(uint(packagesID))
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if packages == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Packages Not Found!"))
		return
	}

	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Invalid request!"))
		return
	}

	validate = validator.New()
	err = validate.Struct(input)

	if err != nil {
		errMap := helpers.ErrorMapValidation(err)
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Bad Request!", gin.H{
			"error": errMap,
		}))
		return
	}

	err = ctl.service.Modify(input, uint(packagesID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Update Packages Success",
		Status:  true,
	})
}

func (ctl *controller) DeletePackages(c *gin.Context) {
	packagesID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	packages, err := ctl.service.FindByID(uint(packagesID))

	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if packages == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Packages Not Found!"))
		return
	}

	err = ctl.service.Remove(uint(packagesID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Delete Packages Success",
		Status:  true,
	})
}
