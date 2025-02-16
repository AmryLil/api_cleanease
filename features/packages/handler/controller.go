package handler

import (
	"api_cleanease/features/packages"
	"api_cleanease/features/packages/dtos"
	"api_cleanease/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	service packages.Usecase
}

func New(service packages.Usecase) packages.Handler {
	return &controller{
		service: service,
	}
}

var validate *validator.Validate

func (ctl *controller) GetPackagess(c *gin.Context) {
	var pagination dtos.Pagination
	if err := c.ShouldBindJSON(&pagination); err != nil {
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

func (ctl *controller) CreatePackages(c *gin.Context) {
	var input []dtos.InputPackages

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Invalid request!"))
		return
	}

	err := ctl.service.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Create Packages Success",
		Status:  true,
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
