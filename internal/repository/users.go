package repository

import (
	"errors"
	"mail-service/internal/db"
	"mail-service/internal/helper"
	"mail-service/internal/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepo struct {
	db              *gorm.DB
	CreateUserInput *CreateUserInput
}

type CreateUserInput struct {
	EncryptedEmail string
	Username       string
	Password       string
	PhoneNumber    string
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		db: db.DB,
	}
}

func (r *UserRepo) SetCreateUserInput(input *CreateUserInput) {
	r.CreateUserInput = input
}

func (r *UserRepo) RegisterNewUser() error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		email, err := helper.CreateCryptor().Decrypt(r.CreateUserInput.EncryptedEmail)
		if err != nil {
			logrus.Error(err)
			return err
		}

		if err := r.IsEmailAlreadyTaken(tx, email); err != nil {
			logrus.Error(err)
			return err
		}

		err = tx.Create(&models.User{
			Email:       r.CreateUserInput.EncryptedEmail,
			Password:    r.CreateUserInput.Password,
			Username:    r.CreateUserInput.Username,
			PhoneNumber: r.CreateUserInput.PhoneNumber,
		}).Error
		if err != nil {
			logrus.Error(err)
			return err
		}

		return nil
	})
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

// IsEmailAlreadyTaken will find the given email. Must be encrypted first
func (r *UserRepo) IsEmailAlreadyTaken(tx *gorm.DB, email string) error {
	var found bool
	err := tx.Model(&models.User{}).
		Select("COUNT (*) > 0").
		Where("email = ?", email).
		Scan(&found).Error
	if err != nil {
		logrus.Error(err)
		return err
	}

	if found {
		return errors.New("email already used")
	}

	return nil
}
