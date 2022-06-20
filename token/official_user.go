package token

import (
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"visitor-management-system/config"
	"visitor-management-system/types"
)

func GenerateOfficialUserTokens(email string, id int, sid int) (signed_token string, signed_refreshtoken string, err error) {

	claims := &types.SignedOfficialUserDetails{
		Email:        email,
		Id:           id,
		SubscriberId: sid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshclaims := &types.SignedOfficialUserDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.GetConfig().SecretKey))
	refresh_token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims).SignedString([]byte(config.GetConfig().SecretKey))

	return token, refresh_token, nil
}
