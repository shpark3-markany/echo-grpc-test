package controllers

import (
	"errors"
	"fmt"
	"local/fin/forms"
	"local/fin/models"
	"local/fin/utils"
	"strings"
)

var db = utils.GetDB()

func InvalidParams(params ...string) string {
	join_params := strings.Join(params, ", ")
	var err_message = fmt.Sprintf("invalid request params %s", join_params)
	return err_message
}

var (
	errInternal = errors.New("internal server err")
	errNotFound = errors.New("record not found")
)

// CRUD
// (Get, List, Create, Update, Delete, Reset)
func Get(id uint64) (*models.UserModel, error) {
	var user *models.UserModel
	if db.Model(&models.UserModel{}).Where("id = ?", id).Find(&user); user.Id == 0 {
		return nil, errNotFound
	}
	return user, nil
}

func List() ([]models.UserModel, error) {
	var users []models.UserModel
	if err := db.Model(&models.UserModel{}).Find(&users).Error; err != nil {
		return nil, errInternal
	}
	return users, nil
}

func Create(user *forms.UserForm) error {
	obj := models.UserModel{
		Email:    user.Email,
		UserName: user.UserName,
		Password: user.Password,
		Age:      user.Age,
		Phone:    user.Phone,
		Address:  user.Address,
	}
	tx := db.Begin()
	if err := tx.Model(&models.UserModel{}).Create(&obj).Error; err != nil {
		return err
	}
	tx.Commit()
	return nil
}

// TODO: lifecicle hook error
// func Update(id uint64, update_user *models.UserModel) error {
// 	var user models.UserModel
// 	db.Model(&models.UserModel{Id: id}).Find(&user)
// 	log.Print(user)
// 	if err := db.Model(&user).Where("id = ?", id).Updates(&update_user).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

func Delete(id uint64) error {
	if err := db.Model(&models.UserModel{}).Delete("id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

func Reset() error {
	var users []models.UserModel
	if err := db.Model(&models.UserModel{}).Find(&users).Error; err != nil {
		return err
	}
	if err := db.Model(&models.UserModel{}).Delete(&users).Error; err != nil {
		return err
	}
	return nil
}

// Batch
// (BSave, BDelete)
func BSave(users []models.UserModel) error {
	tx := db.Begin()
	defer tx.Rollback()
	for _, user := range users {
		if err := tx.Model(&user).Save(&user).Error; err != nil {
			return err
		}
	}
	tx.Commit()
	return nil
}

func BDelete(users []models.UserModel) error {
	tx := db.Begin()
	defer tx.Rollback()
	for _, user := range users {
		if err := tx.Model(&models.UserModel{}).Delete(&user).Error; err != nil {
			continue
		}
	}
	tx.Commit()
	return nil
}
