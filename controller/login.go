package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	//"time"
	"fmt"
	"visitor-management-system/const"
	"visitor-management-system/model"
	"visitor-management-system/repository"
	"visitor-management-system/token"
	"visitor-management-system/types"
	"visitor-management-system/utils"
)

func Login(c echo.Context) (err error) {
	var user = new(types.User)
	var admin = new(model.Subscriber)
	var official_user = new(model.OfficialUser)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, consts.BadRequest)
	}

	if validationerr := validate.Struct(user); validationerr != nil {
		return c.JSON(http.StatusInternalServerError, validationerr.Error())
	}

	//should have access?
	if user.Type == "Admin" {
		admin, err = repository.GetSubscriberByEmail(user.Email)
		if admin.Email == "" || err != nil {
			return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
		}
		fmt.Println(admin)
		if err := utils.VerifyPassword(user.Password, admin.Password); err != nil {
			return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
		}
		token, refresh_token, err := token.GenerateAdminTokens(admin.Email, admin.Id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		admin.Token = token
		admin.RefreshToken = refresh_token

		if err := repository.UpdateSubscriber(admin); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

	} else {
		official_user, err = repository.GetOfficialUserByEmail(user.Email)
		if official_user.Email == "" || err != nil {
			return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
		}
		if err := utils.VerifyPassword(user.Password, official_user.Password); err != nil {
			return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
		}
		token, refresh_token, err := token.GenerateOfficialUserTokens(official_user.Email, official_user.Id, official_user.SubscriberId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		official_user.Token = token
		official_user.RefreshToken = refresh_token

		if err := repository.UpdateOfficialUser(official_user); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

	}

	return c.JSON(http.StatusOK, "successful login")
}
