package types

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type SignedAdminDetails struct {
	Id    int
	Email string
	jwt.StandardClaims
}
