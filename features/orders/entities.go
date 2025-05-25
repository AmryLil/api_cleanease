package orders

import (
	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	UserID     uint
	ServiceID  uint
	PackageID  uint
	IsPickup   bool    // true = dijemput kurir, false = antar langsung
	Weight     float64 // Hanya diisi jika IsIndividual == false
	TotalPrice float64
	Status     string `gorm:"type:varchar(50)"`
	Address    string
	Notes      string
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}

func (Orders) TableName() string {
	return "orders"
}

type OrderItem struct {
	gorm.Model
	OrderID             uint
	IndividualPackageID uint
	Qty                 int
	SubTotal            float64 // Qty * Price
}

func (OrderItem) TableName() string {
	return "order_item"
}
