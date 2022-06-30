package utils

import (
	jwt "github.com/dgrijalva/jwt-go"
	"visitor-management-system/config"
	"visitor-management-system/types"
)

func DecodeToken(usertoken string) (*types.SignedUserDetails, error) {
	var claims = &types.SignedUserDetails{}
	_, err := jwt.ParseWithClaims(usertoken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().SecretKey), nil
	})
	return claims, err
}
