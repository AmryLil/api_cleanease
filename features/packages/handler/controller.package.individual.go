package handler

import (
	"api_cleanease/features/packages/dtos"
	"api_cleanease/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (ctl *controller) GetIndividualPackages(c *gin.Context) {
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

	IndividualPackages, total, err := ctl.service.FindAllIndividualPackages(page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if IndividualPackages == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("There is No IndividualPackages!"))
		return
	}

	paginationData := helpers.PaginationResponse(page, pageSize, int(total))

	c.JSON(http.StatusOK, helpers.ResponseGetAllSuccess{
		Status:     true,
		Message:    "Get All IndividualPackages Success",
		Data:       IndividualPackages,
		Pagination: paginationData,
	})
}

func (ctl *controller) IndividualPackagesDetails(c *gin.Context) {
	IndividualPackagesID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	IndividualPackages, err := ctl.service.FindIndividualPackagesByID(uint(IndividualPackagesID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if IndividualPackages == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("IndividualPackages Not Found!"))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseGetDetailSuccess{
		Data:    IndividualPackages,
		Status:  true,
		Message: " Get IndividualPackages Detail Success",
	})
}

func (ctl *controller) CreateIndividualPackages(c *gin.Context) {
	var input []dtos.InputIndividualPackages

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	err := ctl.service.CreateIndividualPackages(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Create IndividualPackages Success",
		Status:  true,
	})
}

func (ctl *controller) UpdateIndividualPackages(c *gin.Context) {
	var input dtos.InputIndividualPackages
	IndividualPackagesID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	IndividualPackages, err := ctl.service.FindIndividualPackagesByID(uint(IndividualPackagesID))
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if IndividualPackages == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("IndividualPackages Not Found!"))
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

	err = ctl.service.ModifyIndividualPackages(input, uint(IndividualPackagesID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Update IndividualPackages Success",
		Status:  true,
	})
}

func (ctl *controller) DeleteIndividualPackages(c *gin.Context) {
	IndividualPackagesID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	IndividualPackages, err := ctl.service.FindIndividualPackagesByID(uint(IndividualPackagesID))

	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if IndividualPackages == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("IndividualPackages Not Found!"))
		return
	}

	err = ctl.service.RemoveIndividualPackages(uint(IndividualPackagesID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Delete IndividualPackages Success",
		Status:  true,
	})
}
