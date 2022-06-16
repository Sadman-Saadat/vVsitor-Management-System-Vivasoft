package helpers

import (
	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"visitor-management-system/config"
	"visitor-management-system/types"
)

func VerifyToken(usertoken string) (bool, error) {
	claims := &types.SignedAdminDetails{}
	flag := false
	token, err := jwt.ParseWithClaims(usertoken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().SecretKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {

			return flag, err
		}

		return flag, err
	}
	if !token.Valid {

		return flag, err
	}
	flag = true

	return flag, nil

}
