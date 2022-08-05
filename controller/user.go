package controller

import (
	"fmt"
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
	tokens.UserName = model_user.Name

	if err := utils.VerifyPassword(user.Password, model_user.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	token, refresh_token, err := token.GenerateUserTokens(model_user.Email, model_user.Id, model_user.UserType, model_user.CompanyId, model_user.BranchId, model_user.SubDomain)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	tokens.User_Token = token
	tokens.User_Refreshtoken = refresh_token
	branch_ids, err := repository.GetBranchRelation(model_user.Id, model_user.CompanyId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(branch_ids)
	branch_details, err := repository.GetBranchList(branch_ids, model_user.CompanyId)
	fmt.Println(branch_details)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	tokens.Branch = branch_details
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
	//var relation = new(model.UserBranchRelation)

	if err := c.Bind(new_user); err != nil {
		return c.JSON(http.StatusBadRequest, consts.BadRequest)
	}

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	// res, err := repository.GetBranchDetails(claims.CompanyId, new_user.BranchId)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, err.Error())
	// }

	//user.BranchId = new_user.BranchId
	//user.Branch.BranchName = res.BranchName
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
	res, err := repository.GetUserByEmail(user.Email, claims.SubDomain)
	if err != nil || res.Email != "" {
		return c.JSON(http.StatusInternalServerError, "already email exists")
	}

	if claims.UserType == "Admin" {
		user_resp, err := repository.CreateUser(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		for _, v := range new_user.BranchId {
			var relation = new(model.UserBranchRelation)
			relation.BranchId = v
			relation.CompanyId = user_resp.CompanyId
			relation.UserId = user_resp.Id
			err := repository.CreateNewUserBranchRelation(relation)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}

		}
	} else {
		return c.JSON(http.StatusUnauthorized, "need admin access")
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
	id, _ := strconv.Atoi(c.Param("id"))
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	if claims.UserType != "Admin" {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}
	//fmt.Println(user.Id)

	if err := repository.DeleteOfficialUser(id); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
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
	//sub_domain := c.QueryParam("sub_domain")

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
	user, err := repository.GetUserByEmail(claims.Email, claims.SubDomain)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "data not found")
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

func GetUserBranchDetails(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	branch_ids, err := repository.GetBranchRelation(id, claims.CompanyId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(branch_ids)
	branch_details, err := repository.GetBranchList(branch_ids, claims.CompanyId)
	fmt.Println(branch_details)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, branch_details)

}
