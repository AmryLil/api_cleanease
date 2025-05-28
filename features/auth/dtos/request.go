package dtos

// InputUser represents user input data for creation and updates
type InputUser struct {
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"password123"`
	UserType int    `json:"user_type" example:"1"`
}

// Pagination represents pagination query parameters
type Pagination struct {
	Page int `query:"page" example:"1"`
	Size int `query:"page_size" example:"5"`
}

// RegisterRequest represents user registration data
type RegisterRequest struct {
	Name     string `json:"name" binding:"required" example:"John Doe"`
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required,min=6" example:"password123"`
}

// LoginRequest represents login credentials
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required" example:"password123"`
}
