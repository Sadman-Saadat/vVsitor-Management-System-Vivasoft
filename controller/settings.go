package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	//"strconv"
	"strings"
	"visitor-management-system/const"
	"visitor-management-system/model"
	"visitor-management-system/repository"
	"visitor-management-system/utils"
)

// func CreateSettings(c echo.Context) error {
// 	var new_settings = new(model.Setting)
// 	if err := c.Bind(new_settings); err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}
// 	auth_token := c.Request().Header.Get("Authorization")
// 	split_token := strings.Split(auth_token, "Bearer ")
// 	claims, err := utils.DecodeToken(split_token[1])
// 	if err != nil {
// 		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
// 	}
// 	new_settings.CompanyId = claims.CompanyId

// 	if validationerr := validate.Struct(new_settings); validationerr != nil {
// 		return c.JSON(http.StatusInternalServerError, validationerr.Error())
// 	}

// 	if claims.UserType != "Admin" {
// 		return c.JSON(http.StatusUnauthorized, "you need admin access")
// 	}

// 	res, err := repository.CreateNewSettings(new_settings)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, res)
// }

func UpadateSettings(c echo.Context) error {
	var new_settings = new(model.Setting)
	if err := c.Bind(new_settings); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}
	new_settings.CompanyId = claims.CompanyId

	if validationerr := validate.Struct(new_settings); validationerr != nil {
		return c.JSON(http.StatusInternalServerError, validationerr.Error())
	}

	if claims.UserType != "Admin" {
		return c.JSON(http.StatusUnauthorized, "you need admin access")
	}

	err = repository.UpdateSettings(new_settings)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, new_settings)
}

func Setting(c echo.Context) error {
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	if claims.UserType != "Admin" {
		return c.JSON(http.StatusUnauthorized, "you need admin access")
	}

	res, err := repository.Setting(claims.CompanyId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
