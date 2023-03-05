package helpers

import "golang.org/x/crypto/bcrypt"

func Encrypt(plainString string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plainString), bcrypt.DefaultCost)
	return string(bytes), err
}

func CompareEncryption(encryptedString, plainString string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedString), []byte(plainString))
	return err == nil
}
