package dtos

type ResPackages struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	PricePerKg  float64 `json:"price_per_kg"`
	Description string  `json:"description"`
}

type ResIndividualPackages struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	PricePerKg float64 `json:"price_per_kg"`
	Cover      string  ` json:"cover"`
}
