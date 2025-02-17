package dtos

type InputPackages struct {
	Name        string  `json:"name" binding:"required"`
	PricePerKg  float64 `json:"price_per_kg" binding:"required"`
	Description string  `json:"description"`
	Cover       string  ` json:"cover"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"page_size"`
}
