package helper

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"

	"github.com/sirupsen/logrus"
)

func ParseEmailAddress(address string) (string, error) {
	mail, err := mail.ParseAddress(address)

	if err != nil {
		return "", fmt.Errorf("email: %s is an invalid email address", address)
	}

	return mail.Address, nil
}

func ValidateEmailAdressList(list []string) error {
	for _, address := range list {
		_, err := mail.ParseAddress(address)

		if err != nil {
			return fmt.Errorf("email: '%s' is an invalid email address", address)
		}
	}

	return nil
}

func IsStringMatchWithRegexp(str string, pattern string) error {
	re, err := regexp.Compile(pattern)
	if err != nil {
		logrus.Error(err)
		return err
	}

	if !re.MatchString(str) {
		return errors.New("string not match with given regexp pattern")
	}

	return nil
}
