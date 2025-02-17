package usecase

import (
	"api_cleanease/features/packages"
	"api_cleanease/features/packages/dtos"
	"api_cleanease/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

func (svc *service) FindAllIndividualPackages(page, size int) ([]dtos.ResIndividualPackages, int64, error) {
	var packagess []dtos.ResIndividualPackages

	packagessEnt, total, err := svc.model.GetAllIndividualPackages(page, size)
	if err != nil {
		log.Error(err.Error())
		return nil, 0, err
	}

	for _, packages := range packagessEnt {
		var data dtos.ResIndividualPackages

		if err := smapping.FillStruct(&data, smapping.MapFields(packages)); err != nil {
			log.Error(err.Error())
			return nil, 0, err
		}

		packagess = append(packagess, data)
	}

	return packagess, total, nil
}

func (svc *service) FindIndividualPackagesByID(packagesID uint) (*dtos.ResIndividualPackages, error) {
	res := dtos.ResIndividualPackages{}
	packages, err := svc.model.SelectIndividualPackagesByID(packagesID)
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

func (svc *service) CreateIndividualPackages(newPackages []dtos.InputIndividualPackages) error {
	packagesList := []packages.IndividualPackages{}

	for _, input := range newPackages {
		var packageItem packages.IndividualPackages
		err := smapping.FillStruct(&packageItem, smapping.MapFields(input))
		if err != nil {
			log.Error(err.Error())
			return nil
		}

		packageItem.ID = helpers.GenerateID()
		packagesList = append(packagesList, packageItem)

	}

	err := svc.model.InsertIndividualPackages(packagesList)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) ModifyIndividualPackages(packagesData dtos.InputIndividualPackages, packagesID uint) error {
	newPackages := packages.IndividualPackages{}

	err := smapping.FillStruct(&newPackages, smapping.MapFields(packagesData))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newPackages.ID = packagesID
	err = svc.model.UpdateIndividualPackages(newPackages)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) RemoveIndividualPackages(packagesID uint) error {
	err := svc.model.DeleteIndividualPackagesByID(packagesID)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
