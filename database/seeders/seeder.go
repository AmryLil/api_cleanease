package seeders

import "gorm.io/gorm"

func SeedAll(db *gorm.DB) {
	SeedServices(db)
	SeedPackages(db)
	SeedIndividualPackages(db)
}
