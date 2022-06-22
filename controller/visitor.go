package controller

import (
	"fmt"
	//"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path"
	// 	"strings"
	// 	"time"
	// 	"visitor-management-system/const"
	"visitor-management-system/model"
	"visitor-management-system/repository"
	// 	"visitor-management-system/types"
	// 	"visitor-management-system/utils"
)

//register
func CreateVisitor(c echo.Context) error {
	var visitor = new(model.Visitor)

	visitor.Name = c.FormValue("name")
	visitor.Address = c.FormValue("address")
	visitor.CompanyRepresentating = c.FormValue("company_rep")
	visitor.Email = c.FormValue("email")
	visitor.Phone = c.FormValue("phone")
	visitor.Arrived = "Yes"
	file, err := c.FormFile("image")

	if file != nil {
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, file.Header)
		}
		defer src.Close()

		uploadedfilename := file.Filename
		uploadedfilepath := path.Join("./images", uploadedfilename)
		fmt.Println(uploadedfilepath)
		dst, err := os.Create(uploadedfilepath)
		defer dst.Close()
		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		if _, err = io.Copy(dst, src); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		visitor.ImageName = uploadedfilename
		visitor.ImagePath = uploadedfilepath

	}

	if err := repository.CreateVisitor(visitor); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, visitor)
}

//get all the visitor
func GetAllVisitor(c echo.Context) error {
	// var visitor = new(model.Visitor)
	// c.Bind(visitor)
	res, err := repository.GetAllVisitor()
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

//details
func GetVisitorDetails(c echo.Context) error {
	var visitor = new(model.Visitor)
	if err := c.Bind(visitor); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res, err := repository.GetVisitorDetails(visitor)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func UpdateVisitor(c echo.Context) error {
	var visitor = new(model.Visitor)
	if err := c.Bind(visitor); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err := repository.UpdateVisitor(visitor)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, "update successful")
}

func GetVisitor(c echo.Context) error {
	var visitor = new(model.Visitor)
	if err := c.Bind(visitor); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res, err := repository.GetVisitor(visitor)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

//func GetVisitorLog(c echo.Context) err
