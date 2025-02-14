package orders

import (
	"api_cleanease/features/auth"
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	CustomerID uint      `gorm:"not null"`
	Customer   auth.User `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	TotalPrice    float64 `gorm:"not null"`
	Status        string  `gorm:"type:varchar(50);default:'pending'"`
	PickupDate    time.Time
	DeliveryDate  *time.Time    `gorm:"default:null"`
	PaymentMethod string        `gorm:"type:varchar(50);default:null"`
	OrderDetails  []OrderDetail `gorm:"foreignKey:OrderID"`
}

func (Orders) TableName() string {
	return "orders"
}

type OrderDetail struct {
	ID         uint    `gorm:"primaryKey"`
	OrderID    uint    `gorm:"not null"`
	ServiceID  uint    `gorm:"not null"`
	PackageID  uint    `gorm:"not null"`
	Quantity   int     `gorm:"not null"`
	Price      float64 `gorm:"not null"`
	TotalPrice float64 `gorm:"not null"`
}

func (OrderDetail) TableName() string {
	return "order_details"
}
