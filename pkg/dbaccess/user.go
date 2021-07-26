package dbaccess

import (
	"errors"
	"fmt"
	"log"

	"github.com/Tak1za/mixr/pkg/models"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(id string) (*models.UserModel, error)
	CreateUser(user *models.UserModel) (string, error)
	UpdateUser(user *models.UserModel, id string) error
	DeleteUser(id string) error
}

func (e *Env) GetUser(id string) (*models.UserModel, error) {
	var dbUser models.UserModel
	if err := e.DB.Model(&models.UserModel{}).Where("id = ?", id).First(&dbUser).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			log.Println(err.Error())
			return nil, fmt.Errorf("no user found with id: %s", id)
		default:
			log.Println(err.Error())
			return nil, fmt.Errorf("error occured while getting user with id: %s", id)
		}
	}

	return &dbUser, nil
}

func (e *Env) CreateUser(user *models.UserModel) (string, error) {
	if err := e.DB.Model(&models.UserModel{}).Create(user).Error; err != nil {
		switch err.(*pgconn.PgError).Code {
		case "23505":
			log.Println(err.Error())
			return "", errors.New("user with the email id is already registered")
		default:
			log.Println(err.Error())
			return "", errors.New("error occured while creating user")
		}
	}

	return user.ID, nil
}

func (e *Env) UpdateUser(user *models.UserModel, id string) error {
	res := e.DB.Model(&models.UserModel{}).Where("id = ?", id).Updates(user)

	if res.RowsAffected == 0 {
		log.Printf("no user found with id: %s\n", id)
		return fmt.Errorf("no user found with id: %s", id)
	}

	if err := res.Error; err != nil {
		log.Println(err.Error())
		return fmt.Errorf("error updating user with id: %s", id)
	}

	return nil
}

func (e *Env) DeleteUser(id string) error {
	res := e.DB.Model(&models.UserModel{}).Where("id = ?", id).Delete(&models.UserModel{})

	if res.RowsAffected == 0 {
		log.Printf("no user found with id: %s\n", id)
		return fmt.Errorf("no user found with id: %s", id)
	}

	if err := res.Error; err != nil {
		log.Println(err.Error())
		return fmt.Errorf("error deleting user with id: %s", id)
	}

	return nil
}
