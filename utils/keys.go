package utils

import (
	"crypto/sha256"
)

func GenerateKey(password string) []byte {
	hashedPassword := sha256.Sum256([]byte(password))
	return hashedPassword[:]
}
