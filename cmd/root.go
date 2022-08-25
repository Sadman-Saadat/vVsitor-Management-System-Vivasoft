/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/labstack/echo/v4"
	"os"
	"visitor-management-system/config"
	"visitor-management-system/database"
	"visitor-management-system/middleware"

	"visitor-management-system/routes"

	"github.com/spf13/cobra"
	//"github.com/swaggo/echo-swagger"
	//_ "visitor-management-system/docs"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "visitor-management-system",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	e := echo.New()
	database.GetDB()
	database.Migration()

	database.GetDB()
	database.Migration()

	config := config.GetConfig()
	routes.Company(e)
	routes.User(e)
	routes.Visitor(e)
	routes.Branch(e)
	routes.Settings(e)
	routes.MasterAdmin(e)

	dg := e.Group("docs")
	dg.GET("/swagger", echo.WrapHandler(middleware.SwaggerDocs()))
	dg.GET("/redoc", echo.WrapHandler(middleware.ReDocDocs()))
	dg.GET("/rapidoc", echo.WrapHandler(middleware.RapiDocs()))
	e.File("/swagger.yaml", "./swagger.yaml")
	e.Start(":" + config.Port)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.visitor-management-system.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
