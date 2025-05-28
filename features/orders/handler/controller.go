package handler

import (
	"api_cleanease/features/orders"
	"api_cleanease/features/orders/dtos"
	"api_cleanease/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	service orders.Usecase
}

func New(service orders.Usecase) orders.Handler {
	return &controller{
		service: service,
	}
}

var validate *validator.Validate

// GetOrderss godoc
// @Summary Get all orders
// @Description Get all orders with pagination
// @Tags Orders
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1) minimum(1)
// @Param size query int false "Page size" default(5) minimum(1)
// @Success 200 {object} helpers.ResponseGetAllSuccess{data=[]dtos.ResOrders,pagination=helpers.Pagination} "Get all orders success"
// @Failure 400 {object} helpers.ResponseError "Invalid pagination data"
// @Failure 404 {object} helpers.ResponseError "No orders found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /orders [get]
func (ctl *controller) GetOrderss(c *gin.Context) {
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

	orderss, total, err := ctl.service.FindAll(page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if orderss == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("There is No Orderss!"))
		return
	}

	paginationData := helpers.PaginationResponse(page, pageSize, int(total))

	c.JSON(http.StatusOK, helpers.ResponseGetAllSuccess{
		Status:     true,
		Message:    "Get All Orderss Success",
		Data:       orderss,
		Pagination: paginationData,
	})
}

// OrdersDetails godoc
// @Summary Get order details
// @Description Get detailed information of a specific order by ID
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} helpers.ResponseGetDetailSuccess{data=dtos.ResOrders} "Get order detail success"
// @Failure 400 {object} helpers.ResponseError "Invalid order ID"
// @Failure 404 {object} helpers.ResponseError "Order not found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /orders/{id} [get]
func (ctl *controller) OrdersDetails(c *gin.Context) {
	ordersID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	orders, err := ctl.service.FindByID(uint(ordersID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if orders == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Orders Not Found!"))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseGetDetailSuccess{
		Data:    orders,
		Status:  true,
		Message: " Get Orders Detail Success",
	})
}

// CreateOrders godoc
// @Summary Create a new order
// @Description Create a new laundry order
// @Tags Orders
// @Accept json
// @Produce json
// @Param order body dtos.InputOrders true "Order data"
// @Success 200 {object} helpers.ResponseCUDSuccess "Create order success"
// @Failure 400 {object} helpers.ResponseError "Invalid request data"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /orders [post]
func (ctl *controller) CreateOrders(c *gin.Context) {
	var input dtos.InputOrders

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Invalid request!"+err.Error()))
		return
	}

	validate = validator.New()

	err := validate.Struct(input)

	if err != nil {
		errMap := helpers.ErrorMapValidation(err)
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Bad Request!", gin.H{
			"error": errMap,
		}))
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, helpers.BuildErrorResponse(err.Error()))
		return
	}
	input.UserID = uint(userID.(int))

	err = ctl.service.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Create Orders Success",
		Status:  true,
	})
}

// UpdateOrders godoc
// @Summary Update order
// @Description Update an existing order by ID
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param order body dtos.InputOrders true "Order update data"
// @Success 200 {object} helpers.ResponseCUDSuccess "Update order success"
// @Failure 400 {object} helpers.ResponseError "Invalid request data or order ID"
// @Failure 404 {object} helpers.ResponseError "Order not found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /orders/{id} [put]
func (ctl *controller) UpdateOrders(c *gin.Context) {
	var input dtos.InputOrders
	ordersID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	orders, err := ctl.service.FindByID(uint(ordersID))
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if orders == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Orders Not Found!"))
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

	err = ctl.service.Modify(input, uint(ordersID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Update Orders Success",
		Status:  true,
	})
}

// DeleteOrders godoc
// @Summary Delete order
// @Description Delete a specific order by ID
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} helpers.ResponseCUDSuccess "Delete order success"
// @Failure 400 {object} helpers.ResponseError "Invalid order ID"
// @Failure 404 {object} helpers.ResponseError "Order not found"
// @Failure 500 {object} helpers.ResponseError "Internal server error"
// @Router /orders/{id} [delete]
func (ctl *controller) DeleteOrders(c *gin.Context) {
	ordersID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	orders, err := ctl.service.FindByID(uint(ordersID))

	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if orders == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("Orders Not Found!"))
		return
	}

	err = ctl.service.Remove(uint(ordersID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Delete Orders Success",
		Status:  true,
	})
}
