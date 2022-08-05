package routes

import (
	"github.com/labstack/echo/v4"
	"visitor-management-system/controller"
	//"visitor-management-system/middleware"
)

func Company(e *echo.Echo) {
	sub := e.Group("/subscriber")
	sub.POST("/registration", controller.Registration)
	// sub.PATCH("/change-subscription", controller.ChangeSubscription, middleware.Authenticate)
	// sub.DELETE("/", controller.CancelSubscription, middleware.Authenticate)

}
