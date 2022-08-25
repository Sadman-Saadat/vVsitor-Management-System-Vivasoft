package controller

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	"time"
	"visitor-management-system/config"
	"visitor-management-system/types"
	// "visitor-management-system/const"
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

	var relation = new(model.UserBranchRelation)
	//bind
	if err := c.Bind(company); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	company.SubDomain = strings.ToLower(company.CompanyName[0:4])
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
	company.Status = false

	pack, err := repository.GetPackageById(company.Package_Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	company.Subscription_Start = time.Now().Local()
	company.Subscription_End = time.Now().Local().Add(time.Hour * time.Duration(pack.Duration))

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

	// password, err := utils.GenerateRandomPassword()
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }

	// admin.Password, err = utils.Encrypt(password)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }
	//create subscriber
	admin.CompanyId = res.Id
	admin.Name = company.SubscriberName
	admin.Email = company.SubscriberEmail
	admin.UserType = "Admin"
	admin.SubDomain = company.SubDomain
	//admin.BranchId = branchdetails.Id

	if validationerr := validate.Struct(admin); validationerr != nil {
		return c.JSON(http.StatusInternalServerError, validationerr.Error())
	}
	admin_user, err := repository.CreateUser(admin)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	relation.BranchId = branchdetails.Id
	relation.CompanyId = admin_user.CompanyId
	relation.UserId = admin_user.Id

	if err := repository.CreateNewUserBranchRelation(relation); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	full_link := fmt.Sprintf("%s%s", config.GetConfig().Link, utils.EncryptString([]byte("example key 1234"), strconv.Itoa(admin_user.Id)))

	//confirmation mail
	if err := utils.SendEmail(admin.Email, "", admin.SubDomain, full_link, company.CompanyName); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, "Registration successful")
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

// func CancelSubscription(c echo.Context) error {
// 	var subscription = new(model.Subscription)

// 	auth_token := c.Request().Header.Get("Authorization")
// 	split_token := strings.Split(auth_token, "Bearer ")
// 	claims, err := utils.DecodeToken(split_token[1])
// 	if err != nil {
// 		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
// 	}
// 	subscription.CompanyId = claims.CompanyId
// 	subscription.Subscription_type = "cancel"

// 	if claims.UserType != "Admin" {
// 		return c.JSON(http.StatusUnauthorized, "not authorized")
// 	}
// 	subscription.Subscription_start = time.Now().Local()
// 	subscription.Subscription_end = time.Now().Local().Add(time.Hour * time.Duration(720))
// 	if err := repository.CancelSubscription(subscription); err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(http.StatusOK, subscription)
// }

func SetAdminPassword(c echo.Context) error {
	var password = new(types.Password)
	if err := c.Bind(password); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	id := c.QueryParam("token")
	decoded_token := utils.DecryptString([]byte("example key 1234"), id)
	val, _ := strconv.Atoi(decoded_token)
	fmt.Println(password.Password)
	encrypted_password, err := utils.Encrypt(password.Password)
	fmt.Println(encrypted_password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed to encrypt")
	}
	if password.Password == password.ConfirmPassword {
		admin, err := repository.SetAdminPassword(val, encrypted_password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to update user")
		}
		if err := repository.SetCompanyStatus(admin.CompanyId); err != nil {
			return c.JSON(http.StatusInternalServerError, "failed to set Company Status")
		}
		return c.JSON(http.StatusOK, "Verification Successful")
	}

	return c.JSON(http.StatusOK, "Verification not Successful")
}

func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, "healthy")
}

func Healthcheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "ki obostha?")
}
