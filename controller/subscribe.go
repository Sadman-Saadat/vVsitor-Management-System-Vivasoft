package controller

import (
	//"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	//"visitor-management-system/model"
	//	"visitor-management-system/repository"
)

func Subscribe(c echo.Context) error {
	return c.JSON(http.StatusCreated, nil)
}
