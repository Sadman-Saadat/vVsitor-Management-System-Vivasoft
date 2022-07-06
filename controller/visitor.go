package controller

import (
	"fmt"
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
	"visitor-management-system/utils"
)

// swagger:route POST /visitor/registration Visitor CreateVisitor
// Create a new Visitor
// responses:
//	201: Genericsuccess
//	400: ClientError
//	404: ServerError
//	500: ServerError
//     Security:
//     - AuthToken
//
//     Consumes:
//     - multipart/form-data
//
//     Produces:
//     - application/json
//
//     SecurityDefinitions:
//     AuthToken:
//          type: apiKey
//          name: bearer
//          in: header

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

	is_registered, err := repository.IsVistorRegistered(visitor.Email, claims.CompanyId)

	if is_registered != true || err != nil {
		return c.JSON(http.StatusBadRequest, "user already registered")
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

// swagger:route GET /visitor/get-all Visitor Visitors
// All the registered visitor for specific company
// responses:
//	200: AllVisitor
//	400: ClientError
//	404: ServerError
//	500: ServerError
//     Security:
//     - AuthToken
//
//     Consumes:
//     - multipart/form-data
//
//     Produces:
//     - application/json
//
//     SecurityDefinitions:
//     AuthToken:
//          type: apiKey
//          name: bearer
//          in: header

//get all the visitor
func GetAllVisitor(c echo.Context) error {
	// var visitor = new(model.Visitor)
	// c.Bind(visitor)
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}
	res, err := repository.GetAllVisitor(claims.CompanyId)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

// swagger:route GET /visitor/details Visitor AllDetails
// All the visits of a visitor to that company
// responses:
//	200: Alltrackdetaails
//	400: ClientError
//	404: ServerError
//	500: ServerError
//     Security:
//     - AuthToken
//
//     Consumes:
//     - multipart/form-data
//
//     Produces:
//     - application/json
//
//     SecurityDefinitions:
//     AuthToken:
//          type: apiKey
//          name: bearer
//          in: header

//details
func GetVisitorDetails(c echo.Context) error {
	var visitor = new(model.Visitor)
	if err := c.Bind(visitor); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}
	res, err := repository.GetVisitorDetails(visitor, claims.CompanyId)
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
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}
	err = repository.UpdateVisitor(visitor, claims.CompanyId)
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

// swagger:route GET /visitor/search Visitor Isregistered
// is the visitor registered
// responses:
//	302: Visitordetails
//	400: ClientError
//	404: ServerError
//	500: ServerError
//     Security:
//     - AuthToken
//
//     Consumes:
//     - multipart/form-data
//
//     Produces:
//     - application/json
//
//     SecurityDefinitions:
//     AuthToken:
//          type: apiKey
//          name: bearer
//          in: header

func SearchVisitor(c echo.Context) error {
	var visitor = new(model.Visitor)
	if err := c.Bind(visitor); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}
	res, err := repository.Search(visitor, claims.CompanyId)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusFound, res)
}

// swagger:route POST /visitor/checkin Visitor Checkin
//checkin
// responses:
//	200: Genericsuccess
//	400: ClientError
//	404: ServerError
//	500: ServerError
//     Security:
//     - AuthToken
//
//     Consumes:
//     - multipart/form-data
//
//     Produces:
//     - application/json
//
//     SecurityDefinitions:
//     AuthToken:
//          type: apiKey
//          name: bearer
//          in: header

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
	info.Status = "Arrived"
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
	info.CheckIn = time.Now().Local().Format("03:04:05 pm")

	if err := repository.CheckIn(info); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, info)
}

// swagger:route GET /visitor/log Visitor TodaysVisitor
// all the visitor present today
// responses:
//	200: LogResponse
//	400: ClientError
//	404: ServerError
//	500: ServerError
//     Security:
//     - AuthToken
//
//     Consumes:
//     - multipart/form-data
//
//     Produces:
//     - application/json
//
//     SecurityDefinitions:
//     AuthToken:
//          type: apiKey
//          name: bearer
//          in: header

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

// swagger:route POST /visitor/checkout/:id Visitor checkout
//checkout
// responses:
//	200: Genericsuccess
//	400: ClientError
//	404: ServerError
//	500: ServerError
//     Security:
//     - AuthToken
//
//     Consumes:
//     - multipart/form-data
//
//     Produces:
//     - application/json
//
//     SecurityDefinitions:
//     AuthToken:
//          type: apiKey
//          name: bearer
//          in: header
func CheckOut(c echo.Context) error {
	var visitor = new(model.Visitor)
	var record = new(model.Record)
	visitor.Id, _ = strconv.Atoi(c.Param("id"))
	record.VId = visitor.Id
	res, err := repository.GetVisitor(visitor)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	record.VisitorName = res.Name
	record.VisitorEmail = res.Email
	record.VisitorPhone = res.Phone
	record.CompanyRepresentating = res.CompanyRepresentating
	record.CompanyId = res.CompanyId
	record.VisitorAddress = res.Address
	track_res, err := repository.GetTrackDetails(res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	track_res.CheckOut = time.Now().Local().Format("03:04:05 pm")
	track_res.Status = "left"
	if err := repository.CheckOut(res, track_res); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	record.AppointedTo = track_res.AppointedTo
	record.LuggageToken = track_res.LuggageToken
	record.Date = track_res.Date
	record.CheckIn = track_res.CheckIn
	record.CheckOut = track_res.CheckOut

	if err := repository.CreateRecord(record); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, record)
}
