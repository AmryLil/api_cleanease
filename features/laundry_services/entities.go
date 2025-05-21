package laundry_services

import (
	"api_cleanease/features/laundry_packages"

	"gorm.io/gorm"
)

type Services struct {
	gorm.Model

	ID          uint                        `gorm:"primaryKey"`
	Name        string                      `gorm:"type:varchar(255);not null"`
	Description string                      `gorm:"type:text" `
	Packages    []laundry_packages.Packages `gorm:"foreignKey:ServiceID"`
}

func (Services) TableName() string {
	return "services"
}
