package routes

import (
	"github.com/labstack/echo/v4"
	"visitor-management-system/controller"
	//"visitor-management-system/middleware"
)

func MasterAdmin(e *echo.Echo) {
	sub := e.Group("/admin")
	sub.POST("/create", controller.CreateMasterAdmin)
	sub.POST("/login", controller.MasterLogin)
	sub.POST("/package/create", controller.CreatePackage)
	sub.GET("/company-list", controller.CompanyList)
	sub.GET("/package-list", controller.Packagelist)

	// sub.PATCH("/change-subscription", controller.ChangeSubscription, middleware.Authenticate)
	// sub.DELETE("/", controller.CancelSubscription, middleware.Authenticate)

}
