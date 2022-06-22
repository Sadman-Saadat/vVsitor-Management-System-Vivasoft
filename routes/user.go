package routes

import (
	"github.com/labstack/echo/v4"
	"visitor-management-system/controller"
	"visitor-management-system/middleware"
)

func OfficialUser(e *echo.Echo) {
	sub1 := e.Group("/subscriber")
	sub1.POST("/create-user", controller.CreateOfficialUser, middleware.Authenticate)
	sub2 := e.Group("/official-user")
	sub2.GET("/get-all", controller.GetAllOfficialUser, middleware.Authenticate)
	sub2.DELETE("/", controller.DeleteOfficialUser, middleware.Authenticate)
}
