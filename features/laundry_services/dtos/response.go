package dtos

import (
	"api_cleanease/features/laundry_packages/dtos"
)

type ResServices struct {
	ID          uint               `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Packages    []dtos.ResPackages `json:"packages,omitempty"`
}
