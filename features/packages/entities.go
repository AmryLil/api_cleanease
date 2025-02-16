package packages

import (
	"gorm.io/gorm"
)

type Packages struct {
	gorm.Model

	ID          uint    `gorm:"primaryKey"`
	ServiceID   uint    `gorm:"not null" `
	Name        string  `gorm:"type:varchar(255);not null" `
	PricePerKg  float64 `gorm:"type:decimal(10,2);not null" `
	Description string  `gorm:"type:text" `
	Cover       string  `gorm:"type:varchar(255)" `
}

func (Packages) TableName() string {
	return "packages"
}
