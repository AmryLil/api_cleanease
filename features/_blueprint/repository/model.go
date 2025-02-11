package repository

import (
	"api_cleanease/features/_blueprint"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) _blueprint.Repository {
	return &model{
		db: db,
	}
}

func (mdl *model) GetAll(page, size int) ([]_blueprint.Placeholder, int64, error) {
	var placeholders []_blueprint.Placeholder
	var total int64

	if err := mdl.db.Model(&placeholders).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size

	err := mdl.db.Offset(offset).Limit(size).Find(&placeholders).Error

	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	return placeholders, total, nil
}

func (mdl *model) Insert(newPlaceholder _blueprint.Placeholder) error {
	err := mdl.db.Create(&newPlaceholder).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectByID(placeholderID uint) (*_blueprint.Placeholder, error) {
	var placeholder _blueprint.Placeholder
	err := mdl.db.First(&placeholder, placeholderID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &placeholder, nil
}

func (mdl *model) Update(placeholder _blueprint.Placeholder) error {
	err := mdl.db.Updates(&placeholder).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteByID(placeholderID uint) error {
	err := mdl.db.Delete(&_blueprint.Placeholder{}, placeholderID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
