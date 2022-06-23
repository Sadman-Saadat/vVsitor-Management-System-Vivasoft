package utils

import (
	"github.com/sethvargo/go-password/password"
	"strings"
)

func GenerateString() (string, error) {
	res, err := password.Generate(4, 0, 0, false, false)
	return res, err
}

func GenerateFile(name string) string {
	user_name := strings.Split(name, " ")
	random, _ := GenerateString()
	filename := user_name[0] + "_" + random + ".jpg"
	return filename

}
