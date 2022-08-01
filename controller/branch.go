package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	"visitor-management-system/const"
	"visitor-management-system/model"
	"visitor-management-system/repository"
	"visitor-management-system/utils"
)

func CreateBrach(c echo.Context) error {
	var new_branch = new(model.Branch)
	if err := c.Bind(new_branch); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}
	new_branch.CompanyId = claims.CompanyId

	if validationerr := validate.Struct(new_branch); validationerr != nil {
		return c.JSON(http.StatusInternalServerError, validationerr.Error())
	}

	if claims.UserType != "Admin" {
		return c.JSON(http.StatusUnauthorized, "you need admin access")
	}

	resp, err := repository.IsBranchValid(claims.CompanyId, new_branch.BranchName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if resp != 0 {
		return c.JSON(http.StatusInternalServerError, "branch already exists")
	}

	res, err := repository.CreateNewBranch(new_branch)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func UpadateBrach(c echo.Context) error {
	var updated_branch = new(model.Branch)
	if err := c.Bind(updated_branch); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}
	updated_branch.CompanyId = claims.CompanyId

	if validationerr := validate.Struct(updated_branch); validationerr != nil {
		return c.JSON(http.StatusInternalServerError, validationerr.Error())
	}

	if claims.UserType != "Admin" {
		return c.JSON(http.StatusUnauthorized, "you need admin access")
	}

	err = repository.UpdateBranch(updated_branch)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, " branch updated")
}

func BranchList(c echo.Context) error {
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	if claims.UserType != "Admin" {
		return c.JSON(http.StatusUnauthorized, "you need admin access")
	}

	res, err := repository.BranchList(claims.CompanyId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func DeleteBranch(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	if claims.UserType != "Admin" {
		return c.JSON(http.StatusUnauthorized, "you need admin access")
	}

	err = repository.DeleteBranch(claims.CompanyId, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "branch deleted")
}
