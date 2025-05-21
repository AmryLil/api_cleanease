package repository

import (
	"api_cleanease/features/laundry_packages"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) laundry_packages.Repository {
	return &model{
		db: db,
	}
}

func (mdl *model) GetAll(page, size int) ([]laundry_packages.Packages, int64, error) {
	var packagess []laundry_packages.Packages
	var total int64

	if err := mdl.db.Model(&packagess).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size

	err := mdl.db.Offset(offset).Limit(size).Find(&packagess).Error

	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	return packagess, total, nil
}

func (mdl *model) Insert(newPackages laundry_packages.Packages) error {
	err := mdl.db.Create(&newPackages).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectByID(packagesID uint) (*laundry_packages.Packages, error) {
	var packages laundry_packages.Packages
	err := mdl.db.First(&packages, packagesID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &packages, nil
}

func (mdl *model) Update(packages laundry_packages.Packages) error {
	err := mdl.db.Updates(&packages).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteByID(packagesID uint) error {
	err := mdl.db.Delete(&laundry_packages.Packages{}, packagesID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
