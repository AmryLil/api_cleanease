package repository

import (
	"api_cleanease/features/laundry_packages"

	"github.com/labstack/gommon/log"
)

func (mdl *model) GetAllIndividualPackages(page, size int) ([]laundry_packages.IndividualPackages, int64, error) {
	var packagess []laundry_packages.IndividualPackages
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

func (mdl *model) InsertIndividualPackages(newPackages []laundry_packages.IndividualPackages) error {
	err := mdl.db.Create(&newPackages).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectIndividualPackagesByID(packagesID uint) (*laundry_packages.IndividualPackages, error) {
	var packages laundry_packages.IndividualPackages
	err := mdl.db.First(&packages, packagesID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &packages, nil
}

func (mdl *model) UpdateIndividualPackages(packages laundry_packages.IndividualPackages) error {
	err := mdl.db.Updates(&packages).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteIndividualPackagesByID(packagesID uint) error {
	err := mdl.db.Delete(&laundry_packages.IndividualPackages{}, packagesID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
