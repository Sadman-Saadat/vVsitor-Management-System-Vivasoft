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
	sub1.PATCH("/change-password/:id", controller.ChangePassword, middleware.Authenticate)
	sub1.GET("/get-all", controller.GetAllUser, middleware.Authenticate)
	sub1.GET("/details/:id", controller.GetUserBranchDetails, middleware.Authenticate)
	sub1.DELETE("/:id", controller.DeleteOfficialUser, middleware.Authenticate)
	sub1.GET("/dashboard/:id", controller.GetAllData)
}
