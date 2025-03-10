package usecase

import (
	"api_cleanease/features/_blueprint"
	"api_cleanease/features/_blueprint/dtos"
	"api_cleanease/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model _blueprint.Repository
}

func New(model _blueprint.Repository) _blueprint.Usecase {
	return &service{
		model: model,
	}
}

func (svc *service) FindAll(page, size int) ([]dtos.ResPlaceholder, int64, error) {
	var placeholders []dtos.ResPlaceholder

	placeholdersEnt, total, err := svc.model.GetAll(page, size)
	if err != nil {
		log.Error(err.Error())
		return nil, 0, err
	}

	for _, placeholder := range placeholdersEnt {
		var data dtos.ResPlaceholder

		if err := smapping.FillStruct(&data, smapping.MapFields(placeholder)); err != nil {
			log.Error(err.Error())
			return nil, 0, err
		}

		placeholders = append(placeholders, data)
	}

	return placeholders, total, nil
}

func (svc *service) FindByID(placeholderID uint) (*dtos.ResPlaceholder, error) {
	res := dtos.ResPlaceholder{}
	placeholder, err := svc.model.SelectByID(placeholderID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if placeholder == nil {
		return nil, nil
	}

	err = smapping.FillStruct(&res, smapping.MapFields(placeholder))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &res, nil
}

func (svc *service) Create(newPlaceholder dtos.InputPlaceholder) error {
	placeholder := _blueprint.Placeholder{}

	err := smapping.FillStruct(&placeholder, smapping.MapFields(newPlaceholder))
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	placeholder.ID = helpers.GenerateID()
	err = svc.model.Insert(placeholder)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Modify(placeholderData dtos.InputPlaceholder, placeholderID uint) error {
	newPlaceholder := _blueprint.Placeholder{}

	err := smapping.FillStruct(&newPlaceholder, smapping.MapFields(placeholderData))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newPlaceholder.ID = placeholderID
	err = svc.model.Update(newPlaceholder)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Remove(placeholderID uint) error {
	err := svc.model.DeleteByID(placeholderID)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
