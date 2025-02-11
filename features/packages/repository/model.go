package repository

import (
	"api_cleanease/features/packages"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) packages.Repository {
	return &model{
		db: db,
	}
}

func (mdl *model) GetAll(page, size int) ([]packages.Packages, int64, error) {
	var packagess []packages.Packages
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

func (mdl *model) Insert(newPackages packages.Packages) error {
	err := mdl.db.Create(&newPackages).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectByID(packagesID uint) (*packages.Packages, error) {
	var packages packages.Packages
	err := mdl.db.First(&packages, packagesID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &packages, nil
}

func (mdl *model) Update(packages packages.Packages) error {
	err := mdl.db.Updates(&packages).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteByID(packagesID uint) error {
	err := mdl.db.Delete(&packages.Packages{}, packagesID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
