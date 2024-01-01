package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func EncryptFile(filename string, key []byte) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("could not create new cipher: %v", err)
	}

	cipherBytes := make([]byte, aes.BlockSize+len(file))
	iv := cipherBytes[:aes.BlockSize]

	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherBytes[aes.BlockSize:], file)

	err = os.WriteFile(filename+".enc", cipherBytes, 0644)
	if err != nil {
		return fmt.Errorf("could not write encrypted file: %v", err)
	}

	return nil
}

func DecryptFile(filename string, key []byte) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("could not create new cipher: %v", err)
	}

	if len(file) < aes.BlockSize {
		return fmt.Errorf("ciphertext too short")
	}

	iv := file[:aes.BlockSize]
	file = file[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(file, file)

	err = os.WriteFile("dec."+(filename)[:len(filename)-4], file, 0644)
	if err != nil {
		return fmt.Errorf("could not write encrypted file: %v", err)
	}

	return nil
}
