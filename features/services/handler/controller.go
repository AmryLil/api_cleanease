package handler

import (
	"api_cleanease/features/services"
	"api_cleanease/features/services/dtos"
	"api_cleanease/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	service services.Usecase
}

func New(service services.Usecase) services.Handler {
	return &controller{
		service: service,
	}
}

var validate *validator.Validate

func (ctl *controller) GetServicess(c *gin.Context) {
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

	servicess, total, err := ctl.service.FindAll(page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if servicess == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("There is No Servicess!"))
		return
	}

	paginationData := helpers.PaginationResponse(page, pageSize, int(total))

	c.JSON(http.StatusOK, helpers.ResponseGetAllSuccess{
		Status:     true,
		Message:    "Get All Servicess Success",
		Data:       servicess,
		Pagination: paginationData,
	})
}

func (ctl *controller) ServicesDetails(c *gin.Context) {
	servicesID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	services, err := ctl.service.FindByID(uint(servicesID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if services == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Services Not Found!"))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseGetDetailSuccess{
		Data:    services,
		Status:  true,
		Message: " Get Services Detail Success",
	})
}

func (ctl *controller) CreateServices(c *gin.Context) {
	var input []dtos.InputServices

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
		Message: " Create Services Success",
		Status:  true,
	})
}

func (ctl *controller) UpdateServices(c *gin.Context) {
	var input dtos.InputServices
	servicesID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	services, err := ctl.service.FindByID(uint(servicesID))
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if services == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Services Not Found!"))
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

	err = ctl.service.Modify(input, uint(servicesID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Update Services Success",
		Status:  true,
	})
}

func (ctl *controller) DeleteServices(c *gin.Context) {
	servicesID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	services, err := ctl.service.FindByID(uint(servicesID))

	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if services == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Services Not Found!"))
		return
	}

	err = ctl.service.Remove(uint(servicesID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Delete Services Success",
		Status:  true,
	})
}
