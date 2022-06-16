package utils

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"visitor-management-system/config"
	"visitor-management-system/types"
	//"time"
)

func DecodeToken(usertoken string) *types.SignedAdminDetails {
	//var claims *types.SignedDetails
	var claims = &types.SignedAdminDetails{}
	token, err := jwt.ParseWithClaims(usertoken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().SecretKey), nil
	})
	// ... error handling
	if err != nil {
		fmt.Println(token)
	}
	// do something with decoded claims
	return claims
}
