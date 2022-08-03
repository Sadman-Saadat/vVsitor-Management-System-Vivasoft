package controller

import (
	//"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
	"visitor-management-system/const"
	"visitor-management-system/model"
	"visitor-management-system/repository"
	"visitor-management-system/utils"
)

var validate = validator.New()

// swagger:route POST /subscriber/registration Subscriber CreateSub
// Create a new subscriber
// responses:
//	201: Genericsuccess
//	400: ClientError
//	404: ServerError
//	500: ServerError
//     Security:
//     - AuthToken
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     SecurityDefinitions:
//     AuthToken:
//          type: apiKey
//          name: bearer
//          in: header

func Registration(c echo.Context) error {
	var company = new(model.Company)
	var admin = new(model.User)
	var branch = new(model.Branch)
	var settings = new(model.Setting)
	//bind
	if err := c.Bind(company); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	//free trial limit
	if company.Subscription.Subscription_type == "free" {
		company.Subscription.Subscription_start = time.Now().Local()
		company.Subscription.Subscription_end = time.Now().Local().Add(time.Hour * time.Duration(240))
	} else {
		company.Subscription.Subscription_start = time.Now().Local()
		company.Subscription.Subscription_end = time.Now().Local().Add(time.Hour * time.Duration(720))

	}

	company.SubDomain = company.CompanyName[0:4]
	//validate info
	if validationerr := validate.Struct(company); validationerr != nil {
		return c.JSON(http.StatusInternalServerError, validationerr.Error())
	}

	res_com, err := repository.IsCompanyValid(company.CompanyName, company.SubDomain)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if res_com != 0 {
		return c.JSON(http.StatusInternalServerError, "company name or subdomain already exists")
	}

	res, err := repository.RegisterCompany(company)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	branch.BranchName = "Head Office"
	branch.CompanyId = res.Id
	branch.Address = company.Address
	branchdetails, err := repository.CreateNewBranch(branch)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	//randrom password generator

	password, err := utils.GenerateRandomPassword()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	admin.Password, err = utils.Encrypt(password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	//create subscriber
	admin.CompanyId = res.Id
	admin.Name = company.SubscriberName
	admin.Email = company.SubscriberEmail
	admin.UserType = "Admin"
	admin.SubDomain = company.SubDomain
	admin.BranchId = branchdetails.Id

	if validationerr := validate.Struct(admin); validationerr != nil {
		return c.JSON(http.StatusInternalServerError, validationerr.Error())
	}

	if err := repository.CreateUser(admin); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	settings.CompanyId = res.Id
	settings.Email = true
	settings.Image = true

	if _, err := repository.CreateNewSettings(settings); err != nil {
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

// swagger:route PATCH /subscriber/change-subscription Subscriber ChangeSub
// change the subscription
// responses:
//	200: Genericsuccess
//	400: ClientError
//	401: UnAuthorized
//	500: ServerError
//     Security:
//     - AuthToken
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     SecurityDefinitions:
//     AuthToken:
//          type: apiKey
//          name: bearer
//          in: header

func ChangeSubscription(c echo.Context) error {
	var subscription = new(model.Subscription)
	if err := c.Bind(subscription); err != nil {
		return c.JSON(http.StatusBadRequest, consts.BadRequest)
	}
	if validationerr := validate.Struct(subscription); validationerr != nil {
		return c.JSON(http.StatusBadRequest, validationerr.Error())
	}

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}
	subscription.CompanyId = claims.CompanyId
	if claims.UserType != "Admin" {
		return c.JSON(http.StatusUnauthorized, "not authorized")
	}
	subscription.Subscription_start = time.Now().Local()
	subscription.Subscription_end = time.Now().Local().Add(time.Hour * time.Duration(720))
	if subscription.Subscription_type == "free" {
		res, err := repository.GetPreviousSubscription(claims.CompanyId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		if res.Subscription_type == "free" {
			return c.JSON(http.StatusBadRequest, "you have already used free trial")
		}
	}
	if err := repository.ChangeSubscription(subscription); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, subscription)
}

// swagger:route DELETE /subscriber/ Subscriber CancelSub
// cancel the subscription
// responses:
//	200: Genericsuccess
//	400: ClientError
//	401: UnAuthorized
//	500: ServerError
//     Security:
//     - AuthToken
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     SecurityDefinitions:
//     AuthToken:
//          type: apiKey
//          name: bearer
//          in: header

func CancelSubscription(c echo.Context) error {
	var subscription = new(model.Subscription)

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}
	subscription.CompanyId = claims.CompanyId
	subscription.Subscription_type = "cancel"

	if claims.UserType != "Admin" {
		return c.JSON(http.StatusUnauthorized, "not authorized")
	}
	subscription.Subscription_start = time.Now().Local()
	subscription.Subscription_end = time.Now().Local().Add(time.Hour * time.Duration(720))
	if err := repository.CancelSubscription(subscription); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, subscription)
}
