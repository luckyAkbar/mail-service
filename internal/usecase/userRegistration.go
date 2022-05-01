package usecase

import (
	"fmt"
	"mail-service/internal/config"
	"mail-service/internal/db"
	"mail-service/internal/helper"
	"mail-service/internal/repository"
	"mail-service/internal/template"
	"net/url"
	"time"

	"github.com/kumparan/go-utils/encryption"
	sib "github.com/sendinblue/APIv3-go-library/lib"
	"github.com/sirupsen/logrus"
)

type UserRegistrationUseCase struct {
	Email                       string
	Username                    string
	Password                    string
	PasswordConfirmation        string
	PhoneNumber                 string
	EmailCodeConfirmation       string
	HashedEmailCodeConfirmation string
	Cryptor                     *encryption.AESCryptor
}

func NewUserRegitrator() *UserRegistrationUseCase {
	return &UserRegistrationUseCase{
		Cryptor: helper.CreateCryptor(),
	}
}

func (u *UserRegistrationUseCase) SetUserRegistrationData(
	email,
	username,
	phoneNumber,
	password,
	passwordConfirmation string,
) {
	u.Email = email
	u.Username = username
	u.Password = password
	u.PasswordConfirmation = passwordConfirmation
	u.PhoneNumber = phoneNumber
}

func (u *UserRegistrationUseCase) Register() error {
	if err := u.validateEmail(); err != nil {
		return err
	}

	if err := u.validatePassword(); err != nil {
		return err
	}

	userInput, err := u.createUserInputData()
	if err != nil {
		logrus.Error(err)
		return err
	}

	u.generateConfirmationCode()

	userRepo := repository.NewUserRepo()
	userRepo.SetCreateUserInput(userInput)
	if err := userRepo.IsEmailAlreadyTaken(db.DB, userRepo.CreateUserInput.EncryptedEmail); err != nil {
		return ErrMailAdressAlreadyUsed
	}

	if err := userRepo.RegisterNewUser(); err != nil {
		logrus.Error(err)
		return err
	}

	go u.sendConfirmationEmail()

	return nil
}

// validatePassword will only pass if user password are
// match between password and confirmation
// at lease 8 chars long
// contain atleast 1 uppercase and 1 special char
func (u *UserRegistrationUseCase) validatePassword() error {
	if u.Password != u.PasswordConfirmation {
		return ErrPasswordMismatch
	}

	if len(u.Password) < config.PASSWORD_LENGTH_MINIMUM {
		return ErrPasswordTooWeak
	}

	err := helper.IsStringMatchWithRegexp(u.Password, config.REGEXP_PATTERN_ATLEAST_ONE_UPPERCASE)
	if err != nil {
		logrus.Error(err)
		return ErrPasswordTooWeak
	}

	err = helper.IsStringMatchWithRegexp(u.Password, config.REGEXP_PATTER_ATLEASE_ONE_SPECIAL_CHAR)
	if err != nil {
		logrus.Error(err)
		return ErrPasswordTooWeak
	}

	return nil
}

// validateEmail will validate if the given email is valid
// and set the valid result to the struct email prop
func (u *UserRegistrationUseCase) validateEmail() error {
	emailAddr, err := helper.ParseEmailAddress(u.Email)
	if err != nil {
		return ErrMailAdressInvalid
	}

	u.Email = emailAddr
	return nil
}

func (u *UserRegistrationUseCase) createUserInputData() (*repository.CreateUserInput, error) {
	encryptedEmail, err := u.Cryptor.Encrypt(u.Email)
	if err != nil {
		logrus.Error(err)
		return &repository.CreateUserInput{}, err
	}

	password, err := helper.HashString(u.Password)
	if err != nil {
		logrus.Error(err)
		return &repository.CreateUserInput{}, err
	}

	encryptedPhone, err := u.Cryptor.Encrypt(u.PhoneNumber)
	if err != nil {
		logrus.Error(err)
		return &repository.CreateUserInput{}, err
	}

	if err := u.generateConfirmationCode(); err != nil {
		logrus.Error(err)
		return &repository.CreateUserInput{}, err
	}

	return &repository.CreateUserInput{
		EncryptedEmail:              encryptedEmail,
		Username:                    u.Username,
		Password:                    password,
		PhoneNumber:                 encryptedPhone,
		HashedEmailCodeConfirmation: u.HashedEmailCodeConfirmation,
	}, nil
}

func (u *UserRegistrationUseCase) generateConfirmationCode() error {
	code := helper.GenerateRandString(config.USER_CODE_CONFIRMATION_LENGTH, time.Now().Unix())
	hashedCode, err := helper.HashString(code)
	if err != nil {
		logrus.Error(err)
		return err
	}

	u.EmailCodeConfirmation = code
	u.HashedEmailCodeConfirmation = hashedCode
	return nil
}

func (u *UserRegistrationUseCase) sendConfirmationEmail() {
	confirmationLink := fmt.Sprintf("%s%s", config.GetEmailConfirmationCode(), url.QueryEscape(u.EmailCodeConfirmation))
	sibHelper := helper.NewSIBHelper()
	emailContent := sibHelper.CreateEmailContent(
		"Email Confirmation",
		"admin@mail.service.luckyakbar.tech",
		template.GenerateHTMLUserConfirmationTemplate(confirmationLink),
		sib.SendSmtpEmailTo{
			Email: u.Email,
			Name:  u.Username,
		},
	)

	if err := sibHelper.SendEmail(emailContent); err != nil {
		logrus.Error(err)
	}
}
