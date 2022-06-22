package routes

import (
	"github.com/labstack/echo/v4"
	"visitor-management-system/controller"
	//"visitor-management-system/middleware"
)

func Company(e *echo.Echo) {
	sub := e.Group("/subscriber")
	sub.POST("/registration", controller.Registration)
	//sub.PATCH("/change-password", controller.ChangePassword, middleware.Authenticate)
	//sub.GET("/get-all", controller.GetAllSubscriber)
	//sub.POST("/create-user", controller.CreateOfficialUser, middleware.Authenticate)
}
