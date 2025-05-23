package handler

import (
	"api_cleanease/features/laundry_services"
	"api_cleanease/features/laundry_services/dtos"
	"api_cleanease/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	service laundry_services.Usecase
}

func New(service laundry_services.Usecase) laundry_services.Handler {
	return &controller{
		service: service,
	}
}

var validate *validator.Validate

// GetServicess godoc
// @Summary Get all laundry services
// @Description Get all laundry services with pagination
// @Tags Services
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1) minimum(1)
// @Param size query int false "Page size" default(5) minimum(1)
// @Success 200 {object} helpers.ResponseGetAllSuccess{data=[]dtos.ResServices,pagination=helpers.PaginationData} "Get all services success"
// @Failure 400 {object} helpers.ResponseError "Invalid pagination data"
// @Failure 404 {object} helpers.ResponseError "No services found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /services [get]
func (ctl *controller) GetServicess(c *gin.Context) {
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

// ServicesDetails godoc
// @Summary Get service details
// @Description Get detailed information of a specific laundry service by ID
// @Tags Services
// @Accept json
// @Produce json
// @Param id path int true "Service ID" minimum(1)
// @Success 200 {object} helpers.ResponseGetDetailSuccess{data=dtosResServicesServicesResponse} "Get service detail success"
// @Failure 400 {object} helpers.ResponseError "Invalid service ID"
// @Failure 404 {object} helpers.ResponseError "Service not found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /services/{id} [get]
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

// CreateServices godoc
// @Summary Create new laundry services
// @Description Create one or multiple new laundry services
// @Tags Services
// @Accept json
// @Produce json
// @Param services body []dtos.InputServices true "Array of service data"
// @Success 200 {object} helpers.ResponseCUDSuccess "Create services success"
// @Failure 400 {object} helpers.ResponseError "Invalid request data"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /services [post]
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

// UpdateServices godoc
// @Summary Update laundry service
// @Description Update an existing laundry service by ID
// @Tags Services
// @Accept json
// @Produce json
// @Param id path int true "Service ID" minimum(1)
// @Param service body dtos.InputServices true "Updated service data"
// @Success 200 {object} helpers.ResponseCUDSuccess "Update service success"
// @Failure 400 {object} helpers.ResponseError "Invalid request data or validation error"
// @Failure 404 {object} helpers.ResponseError "Service not found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /services/{id} [put]
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

// DeleteServices godoc
// @Summary Delete laundry service
// @Description Delete an existing laundry service by ID
// @Tags Services
// @Accept json
// @Produce json
// @Param id path int true "Service ID" minimum(1)
// @Success 200 {object} helpers.ResponseCUDSuccess "Delete service success"
// @Failure 400 {object} helpers.ResponseError "Invalid service ID"
// @Failure 404 {object} helpers.ResponseError "Service not found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /services/{id} [delete]
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
