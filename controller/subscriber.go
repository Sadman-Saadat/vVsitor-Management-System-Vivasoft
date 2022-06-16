package controller

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"visitor-management-system/const"
	"visitor-management-system/model"
	"visitor-management-system/repository"
	"visitor-management-system/utils"
)

var validate = validator.New()

func CreateSubscribe(c echo.Context) error {
	var new_subscriber = new(model.Subscriber)
	//bind
	if err := c.Bind(new_subscriber); err != nil {
		return c.JSON(http.StatusBadRequest, consts.BadRequest)
	}
	//current_time := time.Now().Local().Add(time.Hour * time.Duration(2)).Format("2006-01-02 3:4:5 pm")
	new_subscriber.Subscription_start = time.Now().Local().Format("2006-01-02 3:4:5 pm")
	new_subscriber.Subscription_end = time.Now().Local().Add(time.Hour * time.Duration(720)).Format("2006-01-02 3:4:5 pm")
	//validate info
	if validationerr := validate.Struct(new_subscriber); validationerr != nil {
		return c.JSON(http.StatusInternalServerError, validationerr.Error())
	}
	//randrom password generator

	password, err := utils.GenerateRandomPassword()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	new_subscriber.Password, err = utils.Encrypt(password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	//create subscriber
	if err := repository.CreateSub(new_subscriber); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	//confirmation mail
	new_subscriber.Password = password
	if err := utils.SendSubscriptionEmail(new_subscriber); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, new_subscriber)
}

func GetAllSubscriber(c echo.Context) error {
	all_subscriber, err := repository.GetAllSubscriber()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, all_subscriber)
}

func ChangePassword(c echo.Context) error {
	return c.JSON(http.StatusOK, "changed successfully")
}
