package seeders

import (
	"api_cleanease/features/laundry_services"
	"fmt"

	"gorm.io/gorm"
)

func SeedServices(db *gorm.DB) {
	var count int64
	db.Model(&laundry_services.Services{}).Count(&count)
	if count > 0 {
		fmt.Println("✅ Service sudah disediakan, skip seeding.")
		return
	}

	services := []laundry_services.Services{
		{Name: "Cuci Kering", Description: "Tanpa lipat"},
		{Name: "Cuci Lipat", Description: "Dilipat rapi"},
	}

	if err := db.Create(&services).Error; err != nil {
		fmt.Println("❌ Gagal seed service:", err)
	} else {
		fmt.Println("✅ Seed service berhasil.")
	}
}
