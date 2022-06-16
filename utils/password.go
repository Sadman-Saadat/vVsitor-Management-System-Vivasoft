package utils

import (
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) (string, error) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hashpassword), err
}

func VerifyPassword(givenpassword string, storedpassword string) error {

	err := bcrypt.CompareHashAndPassword([]byte(storedpassword), []byte(givenpassword))

	return err
}

func GenerateRandomPassword() (string, error) {
	res, err := password.Generate(9, 4, 0, false, false)
	return res, err
}
