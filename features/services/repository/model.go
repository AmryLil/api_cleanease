package repository

import (
	"api_cleanease/features/services"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) services.Repository {
	return &model{
		db: db,
	}
}

func (mdl *model) GetAll(page, size int) ([]services.Services, int64, error) {
	var servicess []services.Services
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

func (mdl *model) Insert(newServices services.Services) error {
	err := mdl.db.Create(&newServices).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectByID(servicesID uint) (*services.Services, error) {
	var services services.Services
	err := mdl.db.First(&services, servicesID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &services, nil
}

func (mdl *model) Update(services services.Services) error {
	err := mdl.db.Updates(&services).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteByID(servicesID uint) error {
	err := mdl.db.Delete(&services.Services{}, servicesID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
