package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
	"visitor-management-system/const"
	"visitor-management-system/model"
	"visitor-management-system/repository"
	"visitor-management-system/types"
	"visitor-management-system/utils"
)

var validate = validator.New()

func Registration(c echo.Context) error {
	var company = new(model.Company)
	var admin = new(model.User)
	//bind
	if err := c.Bind(company); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// //current_time := time.Now().Local().Add(time.Hour * time.Duration(2)).Format("2006-01-02 3:4:5 pm")
	company.Subscription.Subscription_start = time.Now().Local()
	company.Subscription.Subscription_end = time.Now().Local().Add(time.Hour * time.Duration(720))
	//validate info
	if validationerr := validate.Struct(company); validationerr != nil {
		return c.JSON(http.StatusInternalServerError, validationerr.Error())
	}

	if err := repository.RegisterCompany(company); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	// //randrom password generator

	password, err := utils.GenerateRandomPassword()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	admin.Password, err = utils.Encrypt(password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	//create subscriber
	admin.CompanyId = company.Id
	admin.Name = company.SubscriberName
	admin.Email = company.SubscriberEmail
	admin.UserType = "Admin"

	if validationerr := validate.Struct(admin); validationerr != nil {
		return c.JSON(http.StatusInternalServerError, validationerr.Error())
	}

	if err := repository.CreateUser(admin); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	//confirmation mail
	if err := utils.SendEmail(admin, password); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, admin)
}

func GetAllSubscriber(c echo.Context) error {
	all_subscriber, err := repository.GetAllSubscriber()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, all_subscriber)
}

func ChangePassword(c echo.Context) error {
	var password = new(types.Password)

	if err := c.Bind(password); err != nil {
		return c.JSON(http.StatusBadRequest, consts.BadRequest)
	}

	if validationerr := validate.Struct(password); validationerr != nil {
		return c.JSON(http.StatusBadRequest, validationerr.Error())
	}
	//token validation

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}
	user, err := repository.GetUserByEmail(claims.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, consts.UnAuthorized)
	}
	//subscriber validation and update password
	if user.Id == claims.Id && password.Password == password.ConfirmPassword {
		user.Password, err = utils.Encrypt(password.ConfirmPassword)
		if err := repository.UpdateUser(user); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}
	return c.JSON(http.StatusOK, user)
}
