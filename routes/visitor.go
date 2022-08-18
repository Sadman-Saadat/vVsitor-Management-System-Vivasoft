package routes

import (
	"github.com/labstack/echo/v4"
	"visitor-management-system/controller"
	//"visitor-management-system/middleware"
)

func Visitor(e *echo.Echo) {
	sub := e.Group("/visitor")
	//sub.Use(middleware.Authenticate)
	sub.POST("/register", controller.CreateVisitor)
	sub.GET("/get-all", controller.GetAllVisitor)
	sub.GET("/details/:id", controller.GetVisitorDetails)
	sub.PATCH("/update", controller.UpdateVisitor)
	sub.GET("/", controller.GetVisitor)
	sub.GET("/search", controller.SearchVisitor)
	sub.POST("/checkin", controller.CheckIn)
	sub.GET("/log", controller.GetTodaysVisitor)
	sub.POST("/check-out/:id", controller.CheckOut)

}
