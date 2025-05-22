package handler

import (
	user "api_cleanease/features/auth"
	"api_cleanease/features/auth/dtos"
	"api_cleanease/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type controller struct {
	service user.Usecase
}

func New(service user.Usecase) user.Handler {
	return &controller{
		service: service,
	}
}

var validate *validator.Validate

// GetUsers godoc
// @Summary      Get all users with pagination
// @Description  Retrieve a paginated list of all users in the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        page      query     int  false  "Page number"      default(1)
// @Param        page_size query     int  false  "Items per page"   default(5)
// @Success      200       {object}  helpers.ResponseGetAllSuccess
// @Failure      400       {object}  helpers.ResponseError
// @Failure      404       {object}  helpers.ResponseError
// @Failure      500       {object}  helpers.ResponseError
// @Router       /users [get]
func (ctl *controller) GetUsers(c *gin.Context) {
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

	users, total, err := ctl.service.FindAll(page, pageSize)

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if users == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("There is No Users!"))
		return
	}

	paginationData := helpers.PaginationResponse(page, pageSize, int(total))

	c.JSON(http.StatusOK, helpers.ResponseGetAllSuccess{
		Status:     true,
		Message:    "Get All Users Success",
		Data:       users,
		Pagination: paginationData,
	})
}

// UserDetails godoc
// @Summary      Get user by ID
// @Description  Retrieve detailed information about a specific user by their ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  helpers.ResponseGetDetailSuccess
// @Failure      400  {object}  helpers.ResponseError
// @Failure      404  {object}  helpers.ResponseError
// @Failure      500  {object}  helpers.ResponseError
// @Router       /users/{id} [get]
func (ctl *controller) UserDetails(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	user, err := ctl.service.FindByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("User Not Found!"))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseGetDetailSuccess{
		Data:    user,
		Status:  true,
		Message: " Get User Detail Success",
	})
}

// CreateUser godoc
// @Summary      Create a new user
// @Description  Create a new user with the provided information
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      dtos.InputUser  true  "User information"
// @Success      200   {object}  helpers.ResponseCUDSuccess
// @Failure      400   {object}  helpers.ResponseError
// @Failure      500   {object}  helpers.ResponseError
// @Router       /users/create [post]
func (ctl *controller) CreateUser(c *gin.Context) {
	var input dtos.InputUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse("Invalid request!"))
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

	err = ctl.service.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Create User Success",
		Status:  true,
	})
}

// UpdateUser godoc
// @Summary      Update an existing user
// @Description  Update user information by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id    path      int             true  "User ID"
// @Param        user  body      dtos.InputUser  true  "Updated user information"
// @Success      200   {object}  helpers.ResponseCUDSuccess
// @Failure      400   {object}  helpers.ResponseError
// @Failure      404   {object}  helpers.ResponseError
// @Failure      500   {object}  helpers.ResponseError
// @Router       /users/{id} [put]
func (ctl *controller) UpdateUser(c *gin.Context) {
	var input dtos.InputUser
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	user, err := ctl.service.FindByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("User Not Found!"))
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

	err = ctl.service.Modify(input, uint(userID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Update User Success",
		Status:  true,
	})
}

// DeleteUser godoc
// @Summary      Delete a user
// @Description  Delete user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  helpers.ResponseCUDSuccess
// @Failure      400  {object}  helpers.ResponseError
// @Failure      404  {object}  helpers.ResponseError
// @Failure      500  {object}  helpers.ResponseError
// @Router       /users/{id} [delete]
func (ctl *controller) DeleteUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}

	user, err := ctl.service.FindByID(uint(userID))

	if err != nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(err.Error()))
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, helpers.BuildErrorResponse("User Not Found!"))
		return
	}

	err = ctl.service.Remove(uint(userID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.ResponseCUDSuccess{
		Message: " Delete User Success",
		Status:  true,
	})
}

// Login godoc
// @Summary      User login
// @Description  Authenticate user with email/username and password
// @Tags         authentication
// @Accept       json
// @Produce      json
// @Param        credentials  body      dtos.LoginRequest  true  "Login credentials"
// @Success      200          {object}  helpers.ResponseAuth
// @Failure      400          {object}  helpers.ResponseError
// @Failure      401          {object}  helpers.ResponseError
// @Router       /auth/login [post]
func (ctl *controller) Login(c *gin.Context) {
	var input dtos.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(err.Error()))
		return
	}
	user, err := ctl.service.Login(input)

	if err != nil {
		c.JSON(http.StatusUnauthorized, helpers.BuildErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, helpers.ResponseAuth{
		Status:  true,
		Message: " User berhasil login",
		Data:    user,
	})

}
