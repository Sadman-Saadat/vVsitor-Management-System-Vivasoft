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
	visitor.BranchId = claims.BranchId

	is_registered, err := repository.IsVistorRegistered(visitor.Email, claims.CompanyId, claims.BranchId)

	if is_registered != true || err != nil {
		return c.JSON(http.StatusBadRequest, "user already registered")
	}

	res, str, err := utils.ValidateSubscription(claims.CompanyId)
	if res != true || err != nil {
		return c.JSON(http.StatusOK, str)
	}

	visitor.CompanyId = claims.CompanyId
	file, err := c.FormFile("image")

	settings, err := repository.Setting(claims.CompanyId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if settings.Image == "yes" && file == nil || settings.Email == "yes" && visitor.Email == "" {
		return c.JSON(http.StatusBadRequest, "image/email is mandatory")
	}

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
	// if err := c.Bind(visitor); err != nil {
	// 	return c.JSON(http.StatusBadRequest, err.Error())
	// }
	visitor.Id, _ = strconv.Atoi(c.Param("id"))
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
	visitor.Phone = c.Param("phone")
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}
	if claims.UserType == "Admin" {
		res, err := repository.Search(visitor, claims.CompanyId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, res)
	}
	res, err := repository.SearchForSpecificBranch(visitor, claims.CompanyId, claims.BranchId)
	if err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, res)
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
	info.AppointedToPhone = c.FormValue("appointed_to_phone")
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
		return c.JSON(http.StatusUnauthorized, str)
	}

	info.CompanyId = claims.CompanyId
	info.BranchId = claims.BranchId
	//save image
	file, err := c.FormFile("image")

	settings, err := repository.Setting(claims.CompanyId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if settings.Image == "yes" && file == nil {
		return c.JSON(http.StatusBadRequest, "image is mandatory")
	}

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

	// t, _ = time.Parse(shortForm, "2013-Feb-03")
	// fmt.Println(t)
	times := time.Now().Local().Format("2006-01-02")
	fmt.Println(times)
	const shortForm = "2006-01-02"
	info.Date, _ = time.Parse(shortForm, times)
	fmt.Println(info.Date)

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
	var start_date, end_date time.Time
	var order string
	const shortForm = "2006-01-02"
	if c.QueryParam("start_date") != "" && c.QueryParam("end_date") != "" && c.QueryParam("order") != "" {
		start_date, _ = time.Parse(shortForm, c.QueryParam("start_date"))
		end_date, _ = time.Parse(shortForm, c.QueryParam("end_date"))
		order = c.QueryParam("order")

	} else {
		start_date, _ = time.Parse(shortForm, time.Now().Local().Format("2006-01-02"))
		end_date, _ = time.Parse(shortForm, time.Now().Local().Format("2006-01-02"))
		order = "DESC"
	}

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}
	join_sql := "SELECT * FROM track_visitors LEFT JOIN visitors ON track_visitors.v_id = visitors.id"
	if claims.UserType == "Admin" {
		sql := fmt.Sprintf("%s WHERE track_visitors.company_id = %d AND track_visitors.date BETWEEN ? AND ? ORDER BY track_visitors.id %s", join_sql, claims.CompanyId, order)
		res, err := repository.GetTodaysVisitor(sql, start_date, end_date)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, res)
	}
	sql := fmt.Sprintf("%s WHERE track_visitors.company_id = %d AND track_visitors.branch_id = %d AND track_visitors.date BETWEEN ? AND ? ORDER BY track_visitors.id %s", join_sql, claims.CompanyId, claims.BranchId, order)
	res, err := repository.GetTodaysVisitor(sql, start_date, end_date)
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
		return c.JSON(http.StatusInternalServerError, "failed fetch visitor")
	}
	record.Name = res.Name
	record.Email = res.Email
	record.Phone = res.Phone
	record.CompanyRepresentating = res.CompanyRepresentating
	record.CompanyId = res.CompanyId
	record.Address = res.Address
	track_res, err := repository.GetTrackDetails(res)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed fetch track")
	}
	track_res.CheckOut = time.Now().Local().Format("03:04:05 pm")
	track_res.Status = "left"
	if err := repository.CheckOut(res, track_res); err != nil {
		return c.JSON(http.StatusInternalServerError, "failed")
	}

	record.AppointedTo = track_res.AppointedTo
	record.LuggageToken = track_res.LuggageToken
	record.Date = track_res.Date
	record.CheckIn = track_res.CheckIn
	record.CheckOut = track_res.CheckOut

	// if err := repository.CreateRecord(record); err != nil {
	// 	return c.JSON(http.StatusInternalServerError, "failed insert db")
	// }

	return c.JSON(http.StatusOK, record)
}
