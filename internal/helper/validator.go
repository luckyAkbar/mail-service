package helper

import (
	"fmt"
	"net/mail"
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
