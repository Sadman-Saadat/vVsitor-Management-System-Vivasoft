package routes

import (
	"github.com/labstack/echo/v4"
	"visitor-management-system/controller"
	//	"visitor-management-system/middleware"
)

func Visitor(e *echo.Echo) {
	sub := e.Group("/visitor")
	sub.POST("/create", controller.CreateVisitor)
	sub.GET("/get-all", controller.GetAllVisitor)

}
