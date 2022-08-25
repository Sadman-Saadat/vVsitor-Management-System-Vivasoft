package routes

import (
	"github.com/labstack/echo/v4"
	"visitor-management-system/controller"
	//"visitor-management-system/middleware"
)

func Company(e *echo.Echo) {
	e.GET("/health-check", controller.Healthcheck)
	sub := e.Group("/subscriber")
	sub.POST("/registration", controller.Registration)
<<<<<<< HEAD
	sub.PATCH("/change-subscription", controller.ChangeSubscription, middleware.Authenticate)
	sub.DELETE("/", controller.CancelSubscription, middleware.Authenticate)
=======
	sub.POST("/verify", controller.SetAdminPassword)
	e.GET("/health-check", controller.Health)
	// sub.PATCH("/change-subscription", controller.ChangeSubscription, middleware.Authenticate)
	// sub.DELETE("/", controller.CancelSubscription, middleware.Authenticate)
>>>>>>> dad55e8260aedf1d4ad7f78775d3ad4da2c70dee

}
