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
		fmt.Println("official user")
	}

	return c.JSON(http.StatusOK, admin)
}
