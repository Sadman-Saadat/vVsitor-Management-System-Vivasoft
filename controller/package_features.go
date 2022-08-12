package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	"visitor-management-system/model"
	"visitor-management-system/repository"
	// "visitor-management-system/token"
	// "visitor-management-system/types"
	"visitor-management-system/utils"
)

func CreatePackageFeatures(c echo.Context) error {
	var features = new(model.PackageFeatures)
	if err := c.Bind(features); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if validationerr := validate.Struct(features); validationerr != nil {
		return c.JSON(http.StatusInternalServerError, validationerr.Error())
	}
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil || claims.UserType != "Master Admin" {
		return c.JSON(http.StatusUnauthorized, "not authorized")
	}

	_, err = repository.CreatePackageFeatures(features)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Package Features Created Successfully")
}

func GetPackageFeatures(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("package_id"))
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil || claims.UserType != "Master Admin" {
		return c.JSON(http.StatusUnauthorized, "not authorized")
	}
	resp, err := repository.GetPackageFeature(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}
