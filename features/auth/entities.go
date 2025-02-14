package auth

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserType int    `gorm:"column:user_type"`
	Name     string `gorm:"type:varchar(255);column:name"`
	Email    string `gorm:"uniqueIndex;type:varchar(255);column:email"`
	Password string `gorm:"type:varchar(255);column:password"`
	IsActive bool   `gorm:"default:true;column:is_active"`
}

type UserDetails struct {
	gorm.Model
	UserID      uint    `gorm:"column:user_id;not null"`
	Address     *string `gorm:"type:varchar(255);column:address;null"`
	DateOfBirth *string `gorm:"type:date;column:date_of_birth;null"`
	PhoneNumber *string `gorm:"type:varchar(15);column:phone_number;null"`
	User        User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

func (User) TableName() string {
	return "users"
}

func (UserDetails) TableName() string {
	return "user_details"
}
