package repository

import (
	user "api_cleanease/features/auth"
	"errors"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.Repository {
	return &model{
		db: db,
	}
}

func (mdl *model) GetAll(page, size int) ([]user.User, int64, error) {
	var users []user.User
	var total int64

	if err := mdl.db.Model(&users).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * size

	err := mdl.db.Offset(offset).Limit(size).Find(&users).Error

	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	return users, total, nil
}

func (mdl *model) CheckEmailExist(email string) error {
	var user user.User
	result := mdl.db.Where("email = ?", email).First(&user)

	if result.RowsAffected > 0 {
		return errors.New("email already exists")
	}

	return nil
}

func (mdl *model) GetUserByEmail(email string) (*user.User, error) {
	var user user.User
	result := mdl.db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (mdl *model) Insert(newUser user.User) error {
	err := mdl.db.Create(&newUser).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectByID(userID uint) (*user.User, error) {
	var user user.User
	err := mdl.db.First(&user, userID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &user, nil
}

func (mdl *model) Update(user user.User) error {
	err := mdl.db.Updates(&user).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteByID(userID uint) error {
	err := mdl.db.Delete(&user.User{}, userID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
