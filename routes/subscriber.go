package routes

import (
	"github.com/labstack/echo/v4"
	"visitor-management-system/controller"
	"visitor-management-system/middleware"
)

func Subscriber(e *echo.Echo) {
	sub := e.Group("/subscriber")
	sub.POST("/create", controller.CreateSubscribe)
	sub.PATCH("/change-password", controller.ChangePassword, middleware.Authenticate)
	sub.GET("/get-all", controller.GetAllSubscriber)
	sub.POST("/create-user", controller.CreateOfficialUser, middleware.Authenticate)
}
