package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Subscribe(c echo.Context) error {
	return c.JSON(http.StatusCreated, "subscribed")
}
