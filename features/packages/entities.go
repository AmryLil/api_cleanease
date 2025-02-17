package packages

import "gorm.io/gorm"

type Packages struct {
	gorm.Model
	ID          uint                 `gorm:"primaryKey"`
	Name        string               `gorm:"type:varchar(255);not null"`
	PricePerKg  float64              `gorm:"type:decimal(10,2);not null"`
	Description string               `gorm:"type:text"`
	Cover       string               `gorm:"type:varchar(255)"`
	Individual  []IndividualPackages `gorm:"foreignKey:PackagesID"`
}

func (Packages) TableName() string {
	return "packages"
}

type IndividualPackages struct {
	gorm.Model
	PackagesID uint     `gorm:"not null"`
	Name       string   `gorm:"type:varchar(255);not null"`
	Price      float64  `gorm:"type:decimal(10,2);not null"`
	Cover      string   `gorm:"type:varchar(255)"`
	Package    Packages `gorm:"foreignKey:PackagesID"`
}

func (IndividualPackages) TableName() string {
	return "individual_packages"
}
