package usecase

import (
	"api_cleanease/features/laundry_services"
	"api_cleanease/features/laundry_services/dtos"
	"api_cleanease/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model laundry_services.Repository
}

func New(model laundry_services.Repository) laundry_services.Usecase {
	return &service{
		model: model,
	}
}

func (svc *service) FindAll(page, size int) ([]dtos.ResServices, int64, error) {
	var servicess []dtos.ResServices

	servicessEnt, total, err := svc.model.GetAll(page, size)
	if err != nil {
		log.Error(err.Error())
		return nil, 0, err
	}

	for _, services := range servicessEnt {
		var data dtos.ResServices

		if err := smapping.FillStruct(&data, smapping.MapFields(services)); err != nil {
			log.Error(err.Error())
			return nil, 0, err
		}

		servicess = append(servicess, data)
	}

	return servicess, total, nil
}

func (svc *service) FindByID(servicesID uint) (*dtos.ResServices, error) {
	res := dtos.ResServices{}
	services, err := svc.model.SelectByID(servicesID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if services == nil {
		return nil, nil
	}

	err = smapping.FillStruct(&res, smapping.MapFields(services))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &res, nil
}

func (svc *service) Create(newServices []dtos.InputServices) error {
	var servicesList []laundry_services.Services

	for _, input := range newServices {
		var serviceItem laundry_services.Services
		err := smapping.FillStruct(&serviceItem, smapping.MapFields(input))
		if err != nil {
			log.Error(err.Error())
			return err
		}

		serviceItem.ID = helpers.GenerateID()
		servicesList = append(servicesList, serviceItem)
	}

	err := svc.model.Insert(servicesList)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Modify(servicesData dtos.InputServices, servicesID uint) error {
	newServices := laundry_services.Services{}

	err := smapping.FillStruct(&newServices, smapping.MapFields(servicesData))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newServices.ID = servicesID
	err = svc.model.Update(newServices)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Remove(servicesID uint) error {
	err := svc.model.DeleteByID(servicesID)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
