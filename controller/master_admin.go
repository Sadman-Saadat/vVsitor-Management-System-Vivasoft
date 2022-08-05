package controller

import (
	//"fmt"
	//"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	// "time"
	"strconv"
	"visitor-management-system/model"
	"visitor-management-system/repository"
	"visitor-management-system/token"
	"visitor-management-system/types"
	"visitor-management-system/utils"
)

func CreateMasterAdmin(c echo.Context) error {
	var master = new(model.MasterAdmin)
	if err := c.Bind(master); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	master.UserType = "Master Admin"
	password, err := utils.GenerateRandomPassword()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	master.Password, err = utils.Encrypt(password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if validationerr := validate.Struct(master); validationerr != nil {
		return c.JSON(http.StatusInternalServerError, validationerr.Error())
	}
	res, err := repository.GetMasterAdminByEmail(master.Email)
	if err != nil || res.Email != "" {
		return c.JSON(http.StatusUnauthorized, "email already exists")
	}

	resp, err := repository.CreateMasterAdmin(master)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := utils.SendEmail(resp.Email, password, ""); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

func MasterLogin(c echo.Context) error {
	//var master = new(model.MasterAdmin)
	var tokens = new(types.Mastertoken)
	var credentials = new(types.Master)
	if err := c.Bind(credentials); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := repository.GetMasterAdminByEmail(credentials.Email)
	if err != nil || res.Email == "" {
		return c.JSON(http.StatusUnauthorized, "not a master admin")
	}
	token, refresh_token, err := token.GenerateUserTokens(res.Email, res.Id, res.UserType, 0, 0, "", res.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	tokens.Token = token
	tokens.RefreshToken = refresh_token

	return c.JSON(http.StatusOK, tokens)
}

func CreatePackage(c echo.Context) error {
	var pack = new(types.PackageDetails)
	var packages = new(model.Package)
	if err := c.Bind(pack); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if validationerr := validate.Struct(pack); validationerr != nil {
		return c.JSON(http.StatusInternalServerError, validationerr.Error())
	}

	days, _ := strconv.Atoi(pack.Days)
	duration := days * 24
	packages.Duration = duration
	packages.Subscription_type = pack.PackageType

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil || claims.UserType != "Master Admin" {
		return c.JSON(http.StatusUnauthorized, "not authorized")
	}

	res, err := repository.CreatePackage(packages)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "not a master admin")
	}

	return c.JSON(http.StatusCreated, res)
}

func CompanyList(c echo.Context) error {
	var pagination = new(types.PaginationGetAllCompany)
	var search string
	var page, limit, offset int
	if c.QueryParam("page") == "" && c.QueryParam("limit") == "" {
		page = 1
		limit = 3
	} else {
		page, _ = strconv.Atoi(c.QueryParam("page"))
		limit, _ = strconv.Atoi(c.QueryParam("limit"))
	}
	if c.QueryParam("search") != "" {
		search = c.QueryParam("search")
	}

	offset = (page - 1) * limit

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil || claims.UserType != "Master Admin" {
		return c.JSON(http.StatusUnauthorized, "not authorized")
	}

	res, count, err := repository.GetCompanyList(limit, offset, search)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	pagination.Items = res
	pagination.TotalCount = count
	return c.JSON(http.StatusOK, pagination)
}

func Packagelist(c echo.Context) error {
	//var pagination = new(types.PaginationGetAllPackage)

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil || claims.UserType != "Master Admin" {
		return c.JSON(http.StatusUnauthorized, "not authorized")
	}

	res, err := repository.GetPackageList()
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
