package routes

import (
	"github.com/labstack/echo/v4"
	"visitor-management-system/controller"
	"visitor-management-system/middleware"
)

func User(e *echo.Echo) {
	sub1 := e.Group("/user")
	sub1.POST("/login", controller.Login)

	sub1.POST("/create", controller.CreateUser, middleware.Authenticate)
	sub2 := e.Group("/official-user")
	sub2.GET("/get-all", controller.GetAllOfficialUser, middleware.Authenticate)
	sub2.DELETE("/", controller.DeleteOfficialUser, middleware.Authenticate)
}
