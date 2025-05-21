package repository

import (
	"api_cleanease/features/laundry_services"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) laundry_services.Repository {
	return &model{
		db: db,
	}
}

func (mdl *model) GetAll(page, size int) ([]laundry_services.Services, int64, error) {
	var servicess []laundry_services.Services
	var total int64

	if err := mdl.db.Model(&servicess).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size

	err := mdl.db.Offset(offset).Limit(size).Find(&servicess).Error

	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	return servicess, total, nil
}

func (mdl *model) Insert(newServices []laundry_services.Services) error {
	err := mdl.db.Create(&newServices).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectByID(servicesID uint) (*laundry_services.Services, error) {
	var services laundry_services.Services
	err := mdl.db.First(&services, servicesID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &services, nil
}

func (mdl *model) Update(services laundry_services.Services) error {
	err := mdl.db.Updates(&services).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteByID(servicesID uint) error {
	err := mdl.db.Delete(&laundry_services.Services{}, servicesID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
