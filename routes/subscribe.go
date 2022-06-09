package routes

import (
	"github.com/labstack/echo/v4"
	"visitor-management-system/controller"
)

func Subcription(e *echo.Echo) {
	e.GET("/subscribe", controller.Subscribe)
}
