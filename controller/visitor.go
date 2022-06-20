package controller

import (
	//"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	// 	"strings"
	// 	"time"
	// 	"visitor-management-system/const"
	"visitor-management-system/model"
	"visitor-management-system/repository"
	// 	"visitor-management-system/types"
	// 	"visitor-management-system/utils"
)

func CreateVisitor(c echo.Context) error {
	var visitor = new(model.Visitor)
	c.Bind(visitor)
	if err := repository.CreateVisitor(visitor); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, visitor)
}
func GetAllVisitor(c echo.Context) error {
	// var visitor = new(model.Visitor)
	// c.Bind(visitor)
	res, err := repository.GetAllVisitor()
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}
