package controller

import (
	//"fmt"
	//"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	//	"time"
	"visitor-management-system/const"
	"visitor-management-system/model"
	"visitor-management-system/repository"
	//"visitor-management-system/types"
	"visitor-management-system/utils"
)

func CreateOfficialUser(c echo.Context) error {
	// var new_official_user = new(model.User)

	// if err := c.Bind(new_official_user); err != nil {
	// 	return c.JSON(http.StatusBadRequest, consts.BadRequest)
	// }

	// if validationerr := validate.Struct(new_official_user); validationerr != nil {
	// 	return c.JSON(http.StatusBadRequest, validationerr.Error())
	// }
	// auth_token := c.Request().Header.Get("Authorization")
	// split_token := strings.Split(auth_token, "Bearer ")
	// claims, err := utils.DecodeToken(split_token[1])
	// if err != nil {
	// 	return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	// }
	// fmt.Println(claims)
	// //new_official_user.SubscriberId = claims.Id
	// new_official_user.Password, err = utils.GenerateRandomPassword()
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }

	// // if err := utils.SendOfficialCreatedEmail(new_official_user); err != nil {
	// // 	return c.JSON(http.StatusInternalServerError, err.Error())
	// // }

	// new_official_user.Password, err = utils.Encrypt(new_official_user.Password)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }

	// if err := repository.CreateOfficialUser(new_official_user); err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }

	return c.JSON(http.StatusCreated, nil)
}

func GetAllOfficialUser(c echo.Context) error {
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	res, err := repository.GetAllOfficialUsers(claims.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func DeleteOfficialUser(c echo.Context) error {
	var user = new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, consts.BadRequest)
	}

	if err := repository.DeleteOfficialUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "delete successful")
}
