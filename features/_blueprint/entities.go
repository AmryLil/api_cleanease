package _blueprint

import (
	"gorm.io/gorm"
)

type Placeholder struct {
	gorm.Model

	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}

func (Placeholder) TableName() string {
	return "placeholder"
}
