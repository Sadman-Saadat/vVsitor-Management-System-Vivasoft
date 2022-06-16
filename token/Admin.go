package token

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"visitor-management-system/config"
	"visitor-management-system/types"
)

func GenerateAdminTokens(email string, id int) (signed_token string, signed_refreshtoken string, err error) {

	claims := &types.SignedAdminDetails{
		Email: email,
		Id:    id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshclaims := &types.SignedAdminDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.GetConfig().SecretKey))
	refresh_token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims).SignedString([]byte(config.GetConfig().SecretKey))

	return token, refresh_token, nil
}
