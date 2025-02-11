package dtos

type ResPackages struct {
	ID          uint    `json:"id"`
	ServiceID   uint    `json:"service_id"`
	Name        string  `json:"name"`
	PricePerKg  float64 `json:"price_per_kg"`
	Description string  `json:"description"`
}
