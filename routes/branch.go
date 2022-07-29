package routes

import (
	"github.com/labstack/echo/v4"
	"visitor-management-system/controller"
	"visitor-management-system/middleware"
)

func Branch(e *echo.Echo) {
	sub := e.Group("/branch")
	sub.Use(middleware.Authenticate)
	sub.POST("/create", controller.CreateBrach)
	sub.GET("/get-all", controller.BranchList)
	sub.PATCH("/update", controller.UpadateBrach)
	sub.DELETE("/:id", controller.DeleteBranch)

}
