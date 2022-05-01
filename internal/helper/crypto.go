package helper

import (
	"crypto/aes"
	"mail-service/internal/config"
	"math/rand"
	"strings"

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

func GenerateRandString(resultLength int, int64Seeder int64) string {
	var output strings.Builder
	rand.Seed(int64Seeder)
	for i := 0; i < resultLength; i++ {
		c := rand.Intn(len(config.LOWERCASED_APHABETICAL_CHAR))
		r := config.LOWERCASED_APHABETICAL_CHAR[c]

		output.WriteString(string(r))
	}

	return output.String()
}
