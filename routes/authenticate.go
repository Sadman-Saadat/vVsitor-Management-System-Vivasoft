package routes

import (
	"github.com/labstack/echo/v4"
	"visitor-management-system/controller"
)

func Authenticate(e *echo.Echo) {
	e.POST("/login", controller.Login)
}
