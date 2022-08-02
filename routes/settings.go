package routes

import (
	"github.com/labstack/echo/v4"
	"visitor-management-system/controller"
	"visitor-management-system/middleware"
)

func Settings(e *echo.Echo) {
	sub := e.Group("/settings")
	sub.Use(middleware.Authenticate)
	//sub.POST("/create", controller.CreateSettings)
	sub.GET("/get", controller.Setting)
	sub.PATCH("/update", controller.UpadateSettings)
	// sub.DELETE("/:id", controller.DeleteBranch)

}
