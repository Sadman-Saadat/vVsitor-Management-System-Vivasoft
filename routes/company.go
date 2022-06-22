package routes

import (
	"github.com/labstack/echo/v4"
	"visitor-management-system/controller"
	//"visitor-management-system/middleware"
)

func Company(e *echo.Echo) {
	sub := e.Group("/subscriber")
	sub.POST("/registration", controller.Registration)

}
