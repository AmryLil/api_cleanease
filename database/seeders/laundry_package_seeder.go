package seeders

import (
	"api_cleanease/features/laundry_packages"
	"fmt"

	"gorm.io/gorm"
)

func SeedPackages(db *gorm.DB) {
	var count int64
	db.Model(&laundry_packages.Packages{}).Count(&count)
	if count > 0 {
		fmt.Println("✅ Package sudah tersedia, skip seeding.")
		return
	}

	packages := []laundry_packages.Packages{
		{ServiceID: 1, Name: "Reguler", PricePerKg: 7000, Description: "Reguler 2 hari"},
		{ServiceID: 1, Name: "Express", PricePerKg: 12000, Description: "Express 1 hari"},
		{ServiceID: 2, Name: "Kilat", PricePerKg: 15000, Description: "Lipat langsung"},
	}

	if err := db.Create(&packages).Error; err != nil {
		fmt.Println("❌ Gagal seed packages:", err)
	} else {
		fmt.Println("✅ Seed packages berhasil.")
	}
}
