package handler

import (
	"api_cleanease/features/laundry_packages/dtos"
	"api_cleanease/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// GetIndividualPackages godoc
// @Summary Get all individual packages
// @Description Get all individual packages with pagination
// @Tags Individual Packages
// @Accept json
// @Produce json
// @Param pagination body dtos.Pagination true "Pagination parameters"
// @Success 200 {object} helpers.ResponseGetAllSuccess{data=[]dtos.ResIndividualPackages,pagination=helpers.Pagination} "Get all individual packages success"
// @Failure 400 {object} helpers.ResponseError "Invalid pagination data"
// @Failure 404 {object} helpers.ResponseError "No individual packages found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /individual-packages [post]
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

// IndividualPackagesDetails godoc
// @Summary Get individual package details
// @Description Get detailed information of a specific individual package by ID
// @Tags Individual Packages
// @Accept json
// @Produce json
// @Param id path int true "Individual Package ID"
// @Success 200 {object} helpers.ResponseGetDetailSuccess{data=dtos.ResIndividualPackages} "Get individual package detail success"
// @Failure 400 {object} helpers.ResponseError "Invalid package ID"
// @Failure 404 {object} helpers.ResponseError "Individual package not found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /individual-packages/{id} [get]
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

// CreateIndividualPackages godoc
// @Summary Create individual packages
// @Description Create multiple individual packages in batch
// @Tags Individual Packages
// @Accept json
// @Produce json
// @Param packages body []dtos.InputIndividualPackages true "Array of individual packages data"
// @Success 200 {object} helpers.ResponseCUDSuccess "Create individual packages success"
// @Failure 400 {object} helpers.ResponseError "Invalid request data"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /individual-packages [post]
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

// UpdateIndividualPackages godoc
// @Summary Update individual package
// @Description Update an existing individual package by ID
// @Tags Individual Packages
// @Accept json
// @Produce json
// @Param id path int true "Individual Package ID"
// @Param package body dtos.InputIndividualPackages true "Individual package update data"
// @Success 200 {object} helpers.ResponseCUDSuccess "Update individual package success"
// @Failure 400 {object} helpers.ResponseError "Invalid request data or package ID"
// @Failure 404 {object} helpers.ResponseError "Individual package not found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /individual-packages/{id} [put]
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

// DeleteIndividualPackages godoc
// @Summary Delete individual package
// @Description Delete a specific individual package by ID
// @Tags Individual Packages
// @Accept json
// @Produce json
// @Param id path int true "Individual Package ID"
// @Success 200 {object} helpers.ResponseCUDSuccess "Delete individual package success"
// @Failure 400 {object} helpers.ResponseError "Invalid package ID"
// @Failure 404 {object} helpers.ResponseError "Individual package not found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /individual-packages/{id} [delete]
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
