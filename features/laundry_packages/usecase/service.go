package usecase

import (
	"api_cleanease/features/laundry_packages"
	"api_cleanease/features/laundry_packages/dtos"
	"api_cleanease/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model laundry_packages.Repository
}

func New(model laundry_packages.Repository) laundry_packages.Usecase {
	return &service{
		model: model,
	}
}

func (svc *service) FindAll(page, size int) ([]dtos.ResPackages, int64, error) {
	var packagess []dtos.ResPackages

	packagessEnt, total, err := svc.model.GetAll(page, size)
	if err != nil {
		log.Error(err.Error())
		return nil, 0, err
	}

	for _, packages := range packagessEnt {
		var data dtos.ResPackages

		if err := smapping.FillStruct(&data, smapping.MapFields(packages)); err != nil {
			log.Error(err.Error())
			return nil, 0, err
		}

		packagess = append(packagess, data)
	}

	return packagess, total, nil
}

func (svc *service) FindByID(packagesID uint) (*dtos.ResPackages, error) {
	res := dtos.ResPackages{}
	packages, err := svc.model.SelectByID(packagesID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if packages == nil {
		return nil, nil
	}

	err = smapping.FillStruct(&res, smapping.MapFields(packages))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &res, nil
}

func (svc *service) Create(newPackage dtos.InputPackages) error {
	var packageItem laundry_packages.Packages

	err := smapping.FillStruct(&packageItem, smapping.MapFields(newPackage))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	packageItem.ID = helpers.GenerateID()

	err = svc.model.Insert(packageItem) // Tidak pakai slice lagi
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Modify(packagesData dtos.InputPackages, packagesID uint) error {
	newPackages := laundry_packages.Packages{}

	err := smapping.FillStruct(&newPackages, smapping.MapFields(packagesData))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newPackages.ID = packagesID
	err = svc.model.Update(newPackages)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Remove(packagesID uint) error {
	err := svc.model.DeleteByID(packagesID)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
