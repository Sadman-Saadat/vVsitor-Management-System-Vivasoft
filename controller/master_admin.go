package controller

import (
	//"fmt"
	//"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	"time"
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

	if err := utils.SendEmail(resp.Email, password, "", "", ""); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Master Admin Created Successfully")
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
	token, refresh_token, err := token.GenerateUserTokens(res.Email, res.Id, res.UserType, 0, 0, "", res.Name, "")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	tokens.User_Token = token
	tokens.User_Refreshtoken = refresh_token

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

	_, err = repository.CreatePackage(packages)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "not a master admin")
	}

	return c.JSON(http.StatusCreated, "Package Created Successfully")
}

func CompanyList(c echo.Context) error {
	var pagination = new(types.PaginationGetAllCompany)
	var search, status string
	var page, limit, offset, package_id int
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

	if c.QueryParam("status") != "" {
		status = c.QueryParam("status")
	}
	if c.QueryParam("package_id") != "" {
		package_id, _ = strconv.Atoi(c.QueryParam("package_id"))
	}

	offset = (page - 1) * limit

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil || claims.UserType != "Master Admin" {
		return c.JSON(http.StatusUnauthorized, "not authorized")
	}

	res, count, err := repository.GetCompanyList(limit, offset, search, status, package_id)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	pagination.Items = res
	pagination.TotalCount = count
	return c.JSON(http.StatusOK, pagination)
}

func Packagelist(c echo.Context) error {
	res, err := repository.GetPackageList()
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func GetAllCompanyAdmin(c echo.Context) error {
	var pagination = new(types.PaginationGetAllAdmins)
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

	resp, count, err := repository.GetCompanyAdminlist(limit, offset, search)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	pagination.TotalCount = count
	pagination.Items = resp

	return c.JSON(http.StatusOK, pagination)

}

func DeletePackage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err1, err2 := repository.DeletePackage(id)
	if err1 != nil || err2 != nil {
		return c.JSON(http.StatusInternalServerError, err1.Error())
	}
	return c.JSON(http.StatusOK, "Package Successfully Deleted")
}

func UpdatePackagefeatures(c echo.Context) error {
	var updated_feature = new(model.PackageFeatures)
	if err := c.Bind(updated_feature); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if validationerr := validate.Struct(updated_feature); validationerr != nil {
		return c.JSON(http.StatusInternalServerError, validationerr.Error())
	}
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil || claims.UserType != "Master Admin" {
		return c.JSON(http.StatusUnauthorized, "not authorized")
	}

	err = repository.UpdateFeatures(updated_feature)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Package Features Updated Successfully")

}

func UpdateCompanyStatus(c echo.Context) error {
	var status = new(types.Status)
	if err := c.Bind(status); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil || claims.UserType != "Master Admin" {
		return c.JSON(http.StatusUnauthorized, "not authorized")
	}
	err = repository.UpdateCompanyStatus(status.Id, status.Status)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "Updated Company Status Successfully")
}

func AdminPasswordChange(c echo.Context) error {
	user_id, _ := strconv.Atoi(c.Param("id"))
	password, err := utils.GenerateRandomPassword()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	encrypted_password, err := utils.Encrypt(password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	admin, err := repository.GetUserById(user_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	admin.Password = encrypted_password
	if err := repository.UpdateOfficialUser(admin); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := utils.SendEmail(admin.Email, password, admin.SubDomain, "", ""); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Admin Password Reset Successful")
}

func ChangeSubscription(c echo.Context) error {
	var new_subscription = new(types.ChangeSubscription)
	if err := c.Bind(new_subscription); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if validationerr := validate.Struct(new_subscription); validationerr != nil {
		return c.JSON(http.StatusBadRequest, validationerr.Error())
	}

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil || claims.UserType != "Master Admin" {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}
	pack, err := repository.GetPackageById(new_subscription.PackageId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	start := time.Now().Local()
	end := time.Now().Local().Add(time.Hour * time.Duration(pack.Duration))

	if err := repository.ChangeSubscription(new_subscription, start, end); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Subscription Change Successful")
}

func Updatepackage(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var pack = new(types.PackageDetails)
	var new_package = new(model.Package)
	if err := c.Bind(pack); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	new_package.Id = id
	new_package.Subscription_type = pack.PackageType
	days, _ := strconv.Atoi(pack.Days)
	duration := days * 24
	new_package.Duration = duration

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil || claims.UserType != "Master Admin" {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	if err := repository.UpdatePackage(new_package); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Package Updated Successfully")
}

func GetAllAdminData(c echo.Context) error {
	var data = new(types.AdminDataCount)

	packagelist, _ := repository.GetPackageList()
	for _, v := range packagelist {
		count, err := repository.GetCompanyPerPackage(v.Id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		var p = new(types.AdmindataPackage)
		p.Package_type = v.Subscription_type
		p.Count = count
		data.Package = append(data.Package, p)

	}
	res, err := repository.GetCompanyCount(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
