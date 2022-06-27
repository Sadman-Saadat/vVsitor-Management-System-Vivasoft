package controller

import (
	"fmt"
	//"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"visitor-management-system/const"
	"visitor-management-system/model"
	"visitor-management-system/repository"
	// 	"visitor-management-system/types"
	"visitor-management-system/utils"
)

//register
func CreateVisitor(c echo.Context) error {
	var visitor = new(model.Visitor)

	visitor.Name = c.FormValue("name")
	visitor.Address = c.FormValue("address")
	visitor.CompanyRepresentating = c.FormValue("company_rep")
	visitor.Email = c.FormValue("email")
	visitor.Phone = c.FormValue("phone")

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	res, str, err := utils.ValidateSubscription(claims.CompanyId)
	if res != true || err != nil {
		return c.JSON(http.StatusOK, str)
	}

	visitor.CompanyId = claims.CompanyId
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

		uploadedfilename := utils.GenerateFile(visitor.Name)
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
		return c.JSON(http.StatusInternalServerError, err.Error())
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

func SearchVisitor(c echo.Context) error {
	var visitor = new(model.Visitor)
	if err := c.Bind(visitor); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res, err := repository.Search(visitor)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusFound, res)
}

func CheckIn(c echo.Context) error {
	var info = new(model.TrackVisitor)
	//assign value
	id := c.FormValue("v_id")
	info.VId, _ = strconv.Atoi(id)
	fl_num := c.FormValue("floor_number")
	info.FloorNumber, _ = strconv.Atoi(fl_num)
	info.Purpose = c.FormValue("purpose")
	info.LuggageToken = c.FormValue("luggage_token")
	info.AppointedTo = c.FormValue("appointed_to")
	//get company id from token
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	res, str, err := utils.ValidateSubscription(claims.CompanyId)
	if res != true || err != nil {
		return c.JSON(http.StatusOK, str)
	}

	info.CompanyId = claims.CompanyId
	//save image
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
		//get visitor name for image
		var visitor = new(model.Visitor)
		visitor.Id = info.VId
		res, err := repository.GetVisitor(visitor)
		if err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}

		uploadedfilename := utils.GenerateFile(res.Name)
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

		info.ImagePath = uploadedfilepath

	}
	info.Date = time.Now().Local().Format("2006-01-02")
	info.CheckIn = time.Now().Local().Format("3:4:5 pm")

	if err := repository.CheckIn(info); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, info)
}

func GetTodaysVisitor(c echo.Context) error {
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}
	res, err := repository.GetTodaysVisitor(claims.CompanyId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func CheckOut(c echo.Context) error {
	return c.JSON(http.StatusOK, "checked out")
}
