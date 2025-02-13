package auth

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}

func (User) TableName() string {
	return "user"
}
