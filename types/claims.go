package types

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type SignedUserDetails struct {
	Id        int
	Email     string
	CompanyId int
	UserType  string
	SubDomain string
	BranchId  int
	jwt.StandardClaims
}
