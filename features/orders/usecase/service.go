package usecase

import (
	"api_cleanease/features/orders"
	"api_cleanease/features/orders/dtos"
	"api_cleanease/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model orders.Repository
}

func New(model orders.Repository) orders.Usecase {
	return &service{
		model: model,
	}
}

func (svc *service) FindAll(page, size int) ([]dtos.ResOrders, int64, error) {
	var orderss []dtos.ResOrders

	orderssEnt, total, err := svc.model.GetAll(page, size)
	if err != nil {
		log.Error(err.Error())
		return nil, 0, err
	}

	for _, orders := range orderssEnt {
		var data dtos.ResOrders

		if err := smapping.FillStruct(&data, smapping.MapFields(orders)); err != nil {
			log.Error(err.Error())
			return nil, 0, err
		}

		orderss = append(orderss, data)
	}

	return orderss, total, nil
}

func (svc *service) FindByID(ordersID uint) (*dtos.ResOrders, error) {
	res := dtos.ResOrders{}
	orders, err := svc.model.SelectByID(ordersID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if orders == nil {
		return nil, nil
	}

	err = smapping.FillStruct(&res, smapping.MapFields(orders))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &res, nil
}

func (svc *service) Create(newOrders dtos.InputOrders) error {
	orders := orders.Orders{}

	err := smapping.FillStruct(&orders, smapping.MapFields(newOrders))
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	orders.ID = helpers.GenerateID()
	err = svc.model.Insert(orders)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Modify(ordersData dtos.InputOrders, ordersID uint) error {
	newOrders := orders.Orders{}

	err := smapping.FillStruct(&newOrders, smapping.MapFields(ordersData))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newOrders.ID = ordersID
	err = svc.model.Update(newOrders)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Remove(ordersID uint) error {
	err := svc.model.DeleteByID(ordersID)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
