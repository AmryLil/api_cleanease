package seeders

import (
	"api_cleanease/features/laundry_packages"
	"fmt"

	"gorm.io/gorm"
)

func SeedIndividualPackages(db *gorm.DB) {
	var count int64
	db.Model(&laundry_packages.IndividualPackages{}).Count(&count)
	if count > 0 {
		fmt.Println("✅ IndividualPackage sudah tersedia, skip seeding.")
		return
	}

	items := []laundry_packages.IndividualPackages{
		{Name: "Cuci Sepatu", Price: 20000},
		{Name: "Cuci Tas", Price: 25000},
		{Name: "Cuci Helm", Price: 15000},
	}

	if err := db.Create(&items).Error; err != nil {
		fmt.Println("❌ Gagal seed individual packages:", err)
	} else {
		fmt.Println("✅ Seed individual packages berhasil.")
	}
}
