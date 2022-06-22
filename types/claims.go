package types

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type SignedUserDetails struct {
	Id        int
	Email     string
	CompanyId int
	UserType  string
	jwt.StandardClaims
}
