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
	sub.GET("/admin-list", controller.GetAllCompanyAdmin)
	sub.DELETE("/package/delete/:id", controller.DeletePackage)
	sub.POST("/package-features/create", controller.CreatePackageFeatures)
	sub.GET("/package-features/:package_id", controller.GetPackageFeatures)
	sub.PATCH("/package-features/update", controller.UpdatePackagefeatures)
	sub.PATCH("/company/active-status/update", controller.UpdateCompanyStatus)
	sub.PATCH("/company-admin/reset-password/:id", controller.AdminPasswordChange)
	sub.PATCH("/change-subscription", controller.ChangeSubscription)
	sub.PATCH("/package/update/:id", controller.Updatepackage)
	sub.GET("/dashboard", controller.GetAllAdminData)
	// sub.PATCH("/change-subscription", controller.ChangeSubscription, middleware.Authenticate)
	// sub.DELETE("/", controller.CancelSubscription, middleware.Authenticate)

}
