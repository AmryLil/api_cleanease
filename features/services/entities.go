package services

import (
	"api_cleanease/features/packages"

	"gorm.io/gorm"
)

type Services struct {
	gorm.Model

	ID          uint                `gorm:"primaryKey"`
	Name        string              `gorm:"type:varchar(255);not null"`
	Description string              `gorm:"type:text" `
	Packages    []packages.Packages `gorm:"foreignKey:ServiceID"`
}

func (Services) TableName() string {
	return "services"
}
