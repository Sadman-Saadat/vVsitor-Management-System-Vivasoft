package routes

import (
	"github.com/labstack/echo/v4"
	"visitor-management-system/controller"
	"visitor-management-system/middleware"
)

func Visitor(e *echo.Echo) {
	sub := e.Group("/visitor")
	sub.POST("/register", controller.CreateVisitor)
	sub.GET("/get-all", controller.GetAllVisitor)
	sub.GET("/get-by-id", controller.GetVisitorDetails, middleware.Authenticate)
	sub.PATCH("/update", controller.UpdateVisitor, middleware.Authenticate)
	sub.GET("/", controller.GetVisitor, middleware.Authenticate)
	sub.GET("/search", controller.SearchVisitor, middleware.Authenticate)
	sub.POST("/checkin", controller.CheckIn, middleware.Authenticate)
	//	sub.GET("/log-info", controller.GetVisitorLog, middleware.Authenticate)

}
