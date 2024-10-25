package models

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type UserModel struct {
	Id       uint64 `json:"id" gorm:"primaryKey;not null" csv:"id"`
	Email    string `json:"email" gorm:"unique;not null" csv:"email"`
	UserName string `json:"user_name" csv:"user_name"`
	Password string `json:"password" gorm:"not null" csv:"password"`
	Age      uint64 `json:"age" csv:"age"`
	Phone    string `json:"phone" csv:"phone"`
	Address  string `json:"address" csv:"address"`
}

// *begin transaction
// BeforeSave
// BeforeCreate
func (u *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	if email_err := EmailValidator(u, tx); email_err != nil {
		return email_err
	}
	if password_err := PasswordValidator(u, tx); password_err != nil {
		return password_err
	}
	return
}

// BeforeUpdate
func (u *UserModel) BeforeUpdate(tx *gorm.DB) (err error) {
	if email_err := EmailValidator(u, tx); email_err != nil {
		return email_err
	}
	if password_err := PasswordValidator(u, tx); password_err != nil {
		return password_err
	}
	return
}
// BeforeDelete
func (u *UserModel) BeforeDelete(tx *gorm.DB) (err error) {
	if err := tx.Where("id = ?", u.Id).First(&UserModel{}).Error; err != nil {
		return errors.New("record not found")
	}
	return nil
}
// *save before associations
// *insert into database
// *save after associations
// AfterCreate
// AfterUpdate
// AfterDelete
// AfterFind
// AfterSave
// *commit or rollback transaction

func EmailValidator(u *UserModel, tx *gorm.DB) error {
	if err := tx.Where("email = ?", u.Email).First(&UserModel{}).Error; err == nil {
		return fmt.Errorf("the email '%s' you have entered already exists", u.Email)
	} else if u.Email == "" {
		return errors.New("can't use this email")
	}
	return nil
}

func PasswordValidator(u *UserModel, tx *gorm.DB) error {
	log.Printf("password validation, %d", len(u.Password))
	if len(u.Password) < 8 {
		return errors.New("password length must have inserted over 8 chracter")
	}
	return nil
}