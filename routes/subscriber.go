package routes

import (
	"github.com/labstack/echo/v4"
	"visitor-management-system/controller"
	"visitor-management-system/middleware"
)

func Subcription(e *echo.Echo) {
	sub := e.Group("/subscriber")
	sub.POST("/create", controller.CreateSubscribe)
	sub.PATCH("/change-password", controller.ChangePassword, middleware.Authenticate)
	sub.GET("/get-all", controller.GetAllSubscriber)
}
