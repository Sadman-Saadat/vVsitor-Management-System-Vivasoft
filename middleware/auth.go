package middleware

import (
	"visitor-management-system/helpers"
	//"fmt"
	//jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//prepare token
		authToken := c.Request().Header.Get("Authorization")
		splitToken := strings.Split(authToken, "Bearer ")

		if len(splitToken) != 2 {
			return c.String(http.StatusUnauthorized, "you need to login to access this page")
		}

		reqToken := splitToken[1]
		if reqToken == "" {
			return c.String(http.StatusInternalServerError, "failed to get token")
		}
		//send for verification

		ok, err := helpers.VerifyToken(reqToken)
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}

		if !ok {
			return c.String(http.StatusInternalServerError, err.Error())
		}

		//next

		return next(c)
	}
}
