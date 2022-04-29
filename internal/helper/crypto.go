package helper

import (
	"crypto/aes"
	"mail-service/internal/config"

	"github.com/kumparan/go-utils/encryption"
	"golang.org/x/crypto/bcrypt"
)

func CreateCryptor() *encryption.AESCryptor {
	privateKey := config.PrivateKey()
	ivKey := config.IvKey()

	return encryption.NewAESCryptor(privateKey, ivKey, aes.BlockSize)
}

func HashString(str string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
