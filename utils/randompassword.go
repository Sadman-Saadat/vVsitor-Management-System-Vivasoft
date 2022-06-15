package utils

import (
	"github.com/sethvargo/go-password/password"
)

func GenerateRandomPassword() (string, error) {
	res, err := password.Generate(9, 4, 0, false, false)
	return res, err
}
