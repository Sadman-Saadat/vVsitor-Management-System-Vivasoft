package controller

import (
	// /"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	"visitor-management-system/const"
	"visitor-management-system/model"
	"visitor-management-system/repository"
	"visitor-management-system/token"
	"visitor-management-system/types"
	"visitor-management-system/utils"
)

// swagger:route PATCH /user/login USER LoginDetails
// user login
// responses:
//	200: LoginSuccess
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

func Login(c echo.Context) (err error) {
	var user = new(types.User)
	var model_user = new(model.User)
	var tokens = new(types.Token)
	sub_domain := c.QueryParam("subdomain")

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, consts.BadRequest)
	}

	if validationerr := validate.Struct(user); validationerr != nil {
		return c.JSON(http.StatusInternalServerError, validationerr.Error())
	}

	model_user, err = repository.GetUserByEmail(user.Email, sub_domain)
	if model_user.Email == "" || err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	if err := utils.VerifyPassword(user.Password, model_user.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	token, refresh_token, err := token.GenerateUserTokens(model_user.Email, model_user.Id, model_user.UserType, model_user.CompanyId, model_user.BranchId, model_user.SubDomain)
	tokens.User_Token = token
	tokens.User_Refreshtoken = refresh_token

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tokens)
}

// swagger:route POST /user/create USER CreateUser
// Create a new user
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

func CreateUser(c echo.Context) error {
	var user = new(model.User)
	var new_user = new(types.User)

	if err := c.Bind(new_user); err != nil {
		return c.JSON(http.StatusBadRequest, consts.BadRequest)
	}

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	// res, err := repository.GetBranchDetails(claims.CompanyId, new_user.BranchName)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }

	user.BranchId = new_user.BranchId
	user.Name = new_user.Name
	user.Email = new_user.Email
	user.Password = new_user.Password
	user.UserType = "Official"
	user.SubDomain = claims.SubDomain
	//fmt.Println(user.SubDomain)
	//user.Branch.Address = res.Address
	user.CompanyId = claims.CompanyId
	password, err := utils.GenerateRandomPassword()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	user.Password, err = utils.Encrypt(password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if validationerr := validate.Struct(user); validationerr != nil {
		return c.JSON(http.StatusBadRequest, validationerr.Error())
	}

	if claims.UserType == "Admin" {
		if err := repository.CreateUser(user); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
	}

	if err := utils.SendEmail(user, password); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

// swagger:route GET /user/get-all USER AllUser
// details of all users
// responses:
//	200: UserDetails
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

func GetAllUser(c echo.Context) error {
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	res, err := repository.GetAllUsers(claims.CompanyId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

// swagger:route DELETE /user/ USER DeleteUser
// delete a  user
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

func DeleteOfficialUser(c echo.Context) error {
	var user = new(model.User)
	user.Id, _ = strconv.Atoi(c.Param("id"))
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}
	user.CompanyId = claims.CompanyId

	if claims.UserType == "Admin" {
		if err := repository.DeleteOfficialUser(user); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

	}

	return c.JSON(http.StatusOK, "delete successful")
}

// swagger:route POST /user/change-password USER ChangePassword
// change user password
// responses:
//	200: LoginSuccess
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

func ChangePassword(c echo.Context) error {
	var password = new(types.Password)
	sub_domain := c.QueryParam("subdomain")

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
	user, err := repository.GetUserByEmail(claims.Email, sub_domain)
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
