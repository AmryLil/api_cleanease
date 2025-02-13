package repository

import (
	"api_cleanease/features/orders"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) orders.Repository {
	return &model{
		db: db,
	}
}

func (mdl *model) GetAll(page, size int) ([]orders.Orders, int64, error) {
	var orderss []orders.Orders
	var total int64

	if err := mdl.db.Model(&orderss).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size

	err := mdl.db.Offset(offset).Limit(size).Find(&orderss).Error

	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	return orderss, total, nil
}

func (mdl *model) Insert(newOrders orders.Orders) error {
	err := mdl.db.Create(&newOrders).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectByID(ordersID uint) (*orders.Orders, error) {
	var orders orders.Orders
	err := mdl.db.First(&orders, ordersID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &orders, nil
}

func (mdl *model) Update(orders orders.Orders) error {
	err := mdl.db.Updates(&orders).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteByID(ordersID uint) error {
	err := mdl.db.Delete(&orders.Orders{}, ordersID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
