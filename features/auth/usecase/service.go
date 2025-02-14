package usecase

import (
	user "api_cleanease/features/auth"
	"api_cleanease/features/auth/dtos"
	"api_cleanease/helpers"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model user.Repository
	hash  helpers.HashInterface
}

func New(model user.Repository, hash helpers.HashInterface) user.Usecase {
	return &service{
		model: model,
		hash:  hash,
	}
}

func (svc *service) FindAll(page, size int) ([]dtos.ResUser, int64, error) {
	var users []dtos.ResUser

	usersEnt, total, err := svc.model.GetAll(page, size)
	if err != nil {
		log.Error(err.Error())
		return nil, 0, err
	}

	for _, user := range usersEnt {
		var data dtos.ResUser

		if err := smapping.FillStruct(&data, smapping.MapFields(user)); err != nil {
			log.Error(err.Error())
			return nil, 0, err
		}

		users = append(users, data)
	}

	return users, total, nil
}

func (svc *service) FindByID(userID uint) (*dtos.ResUser, error) {
	res := dtos.ResUser{}
	user, err := svc.model.SelectByID(userID)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	err = smapping.FillStruct(&res, smapping.MapFields(user))
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return &res, nil
}

func (svc *service) Create(newUser dtos.InputUser) error {
	user := user.User{}

	err := smapping.FillStruct(&user, smapping.MapFields(newUser))
	if err != nil {
		log.Error(err.Error())
		return nil
	}

	user.ID = helpers.GenerateID()
	user.Password = svc.hash.HashPassword(newUser.Password)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.IsActive = true
	user.UserType = 1
	err = svc.model.Insert(user)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Modify(userData dtos.InputUser, userID uint) error {
	newUser := user.User{}

	err := smapping.FillStruct(&newUser, smapping.MapFields(userData))
	if err != nil {
		log.Error(err.Error())
		return err
	}

	newUser.ID = userID
	err = svc.model.Update(newUser)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (svc *service) Remove(userID uint) error {
	err := svc.model.DeleteByID(userID)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}
