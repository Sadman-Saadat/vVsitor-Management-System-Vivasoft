package controller

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo/v4"
	"image/jpeg"
	"visitor-management-system/config"
	//"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"visitor-management-system/const"
	"visitor-management-system/model"
	"visitor-management-system/repository"
	"visitor-management-system/types"
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
	var visitor_response = new(types.VisitorRegistration)
	var f *os.File
	visitor.Name = c.FormValue("name")
	visitor.Address = c.FormValue("address")
	visitor.CompanyRepresentating = c.FormValue("company_rep")
	visitor.Email = c.FormValue("email")
	visitor.Phone = c.FormValue("phone")
	id, _ := strconv.Atoi(c.FormValue("branch_id"))
	fmt.Println(visitor.BranchId)
	visitor.BranchId = id

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}
	//visitor.BranchId = claims.BranchId
	resp, err := repository.GetBranchDetails(claims.CompanyId, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	fmt.Println(resp)
	visitor.BranchName = resp.BranchName
	if visitor.Email != "" {
		is_registered, err := repository.IsVistorRegistered(visitor.Email, claims.CompanyId)

		if is_registered != true || err != nil {
			return c.JSON(http.StatusBadRequest, "visitor email already registered")
		}

	}
	valid_phone, err := repository.IsPhoneNumberPresent(visitor.Phone, claims.CompanyId)
	if valid_phone != true || err != nil {
		return c.JSON(http.StatusBadRequest, "visitor phone already registered")
	}
	res, str, err, image_bool := utils.ValidateSubscription(claims.CompanyId)
	if res != true || err != nil {
		return c.JSON(http.StatusInternalServerError, str)
	}

	visitor.CompanyId = claims.CompanyId
	image := c.FormValue("image")

	if image != "" && image_bool == false {
		return c.JSON(http.StatusBadRequest, "please change subscription to add visitor image")
	}

	if image != "" {

		uploadedfilename := utils.GenerateFile(visitor.Name)
		uploadedfilepath := path.Join("images", uploadedfilename)

		coI := strings.Index(string(image), ",")
		rawImage := string(image)[coI+1:]
		unbased, err := base64.StdEncoding.DecodeString(string(rawImage))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		res := bytes.NewReader(unbased)
		jpgI, errJpg := jpeg.Decode(res)
		if errJpg == nil {
			f, _ = os.OpenFile("images"+"/"+uploadedfilename, os.O_WRONLY|os.O_CREATE, 0777)
			jpeg.Encode(f, jpgI, &jpeg.Options{Quality: 100})
			fmt.Println("Jpg created")
		} else {
			fmt.Println(errJpg.Error())
		}
		visitor.ImageName = uploadedfilename
		visitor.ImagePath = fmt.Sprintf("%s%s", config.GetConfig().ImageBaseUri, uploadedfilepath)

	}
	v_resp, err := repository.CreateVisitor(visitor)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	visitor_response.Message = "Visitor Registration Successful"
	visitor_response.Item = v_resp
	return c.JSON(http.StatusOK, visitor_response)
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
	var search string
	//branch_id, _ := strconv.Atoi(c.QueryParam("branch_id"))
	//var branch_id = new(types.BranchIds)
	var Pagination = new(types.PaginationGetAllVisitor)

	var page, limit, offset int
	if c.QueryParam("page") == "" && c.QueryParam("limit") == "" {
		page = 1
		limit = 3
	} else {
		page, _ = strconv.Atoi(c.QueryParam("page"))
		limit, _ = strconv.Atoi(c.QueryParam("limit"))
	}

	if c.QueryParam("search") != "" {
		search = c.QueryParam("search")
	}

	offset = (page - 1) * limit
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	res, count, err := repository.GetAllVisitorSpecific(claims.CompanyId, search, limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	Pagination.TotalCount = count
	Pagination.Items = res

	return c.JSON(http.StatusOK, Pagination)

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
	return c.JSON(http.StatusOK, "Visitor Updated Successfully")
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
	search := c.QueryParam("search")
	//branch_id, _ := strconv.Atoi(c.QueryParam("branch_id"))
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	res, err := repository.SearchForSpecificBranch(visitor, claims.CompanyId, search)
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
	var f *os.File
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
	info.BranchId, _ = strconv.Atoi(c.FormValue("branch_id"))
	//get company id from token
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}
	var image_bool = false

	res, str, err, image_bool := utils.ValidateSubscription(claims.CompanyId)
	if res != true || err != nil {
		return c.JSON(http.StatusUnauthorized, str)
	}

	info.CompanyId = claims.CompanyId
	//info.BranchId = claims.BranchId
	//save image
	image := c.FormValue("image")

	if image != "" && image_bool == false {
		return c.JSON(http.StatusBadRequest, "please change subscription to add visitor image")
	}

	if image != "" {
		var visitor = new(model.Visitor)
		visitor.Id = info.VId
		resp, err := repository.GetVisitor(visitor)
		if err != nil {
			return c.JSON(http.StatusOK, err.Error())
		}

		uploadedfilename := utils.GenerateFile(resp.Name)
		uploadedfilepath := path.Join("images", uploadedfilename)

		coI := strings.Index(string(image), ",")
		rawImage := string(image)[coI+1:]
		unbased, err := base64.StdEncoding.DecodeString(string(rawImage))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		res := bytes.NewReader(unbased)
		jpgI, errJpg := jpeg.Decode(res)
		if errJpg == nil {
			f, _ = os.OpenFile("images"+"/"+uploadedfilename, os.O_WRONLY|os.O_CREATE, 0777)
			jpeg.Encode(f, jpgI, &jpeg.Options{Quality: 75})
			fmt.Println("Jpg created")
		} else {
			fmt.Println(errJpg.Error())
		}
		info.ImagePath = fmt.Sprintf("%s%s", config.GetConfig().ImageBaseUri, uploadedfilepath)

	}

	times := time.Now().Local().Format("2006-01-02")

	const shortForm = "2006-01-02"
	info.Date, _ = time.Parse(shortForm, times)

	info.CheckIn = time.Now().Local().Format("03:04:05 pm")

	resdetails, err := repository.GetBranchDetails(claims.CompanyId, info.BranchId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	info.BranchName = resdetails.BranchName

	if err := repository.CheckIn(info); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Visitor Checkin Successful")
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
	var pagination = new(types.PaginationGetAllRecord)
	var order, status, search string
	var branch_id int
	var frequent bool

	const shortForm = "2006-01-02"
	if c.QueryParam("start_date") != "" && c.QueryParam("end_date") != "" && c.QueryParam("order") != "" || c.QueryParam("status") != "" || c.QueryParam("branch_id") != "" || c.QueryParam("search") != "" || c.QueryParam("frequent") != "" {
		start_date, _ = time.Parse(shortForm, c.QueryParam("start_date"))
		end_date, _ = time.Parse(shortForm, c.QueryParam("end_date"))
		order = c.QueryParam("order")
		search = c.QueryParam("search")
		status = c.QueryParam("status")
		branch_id, _ = strconv.Atoi(c.QueryParam("branch_id"))
		frequent, _ = strconv.ParseBool(c.QueryParam("frequent"))

	} else {
		start_date, _ = time.Parse(shortForm, time.Now().Local().Format("2006-01-02"))
		end_date, _ = time.Parse(shortForm, time.Now().Local().Format("2006-01-02"))
		order = "DESC"
		frequent = false
	}
	var page, limit, offset int
	if c.QueryParam("page") == "" && c.QueryParam("limit") == "" {
		page = 1
		limit = 3
	} else {
		page, _ = strconv.Atoi(c.QueryParam("page"))
		limit, _ = strconv.Atoi(c.QueryParam("limit"))
	}

	offset = (page - 1) * limit
	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	res, count, err := repository.GetTodaysVisitor(claims.CompanyId, branch_id, start_date, end_date, status, search, order, offset, limit, frequent)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	pagination.Items = res
	pagination.TotalCount = count

	return c.JSON(http.StatusOK, pagination)
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
	//var visitor = new(model.Visitor)
	//var record = new(model.TrackVisitor)
	id, _ := strconv.Atoi(c.Param("id"))

	auth_token := c.Request().Header.Get("Authorization")
	split_token := strings.Split(auth_token, "Bearer ")
	claims, err := utils.DecodeToken(split_token[1])
	if err != nil {
		return c.JSON(http.StatusUnauthorized, consts.UnAuthorized)
	}

	track_res, err := repository.GetTrackDetails(claims.CompanyId, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "failed fetch track")
	}

	track_res.CheckOut = time.Now().Local().Format("03:04:05 pm")
	track_res.Status = "left"
	if err := repository.CheckOut(id, claims.CompanyId, track_res); err != nil {
		return c.JSON(http.StatusInternalServerError, "failed")
	}

	return c.JSON(http.StatusOK, "Visitor Checkout Successful")
}

func GetImageByPath(c echo.Context) error {
	path := c.QueryParam("path")
	return c.File(path)
}
