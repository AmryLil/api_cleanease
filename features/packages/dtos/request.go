package dtos

type InputPackages struct {
	ServiceID   uint    `json:"service_id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	PricePerKg  float64 `json:"price_per_kg" binding:"required"`
	Description string  `json:"description"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"page_size"`
}
