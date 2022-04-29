package helper

import (
	"crypto/aes"
	"mail-service/internal/config"

	"github.com/kumparan/go-utils/encryption"
)

func CreateCryptor() *encryption.AESCryptor {
	privateKey := config.PrivateKey()
	ivKey := config.IvKey()

	return encryption.NewAESCryptor(privateKey, ivKey, aes.BlockSize)
}
