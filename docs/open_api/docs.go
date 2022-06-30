// Package classification VMS API.
//
// the purpose of this service is to provide & store all visitors of a company and their visit histories
//
//     Schemes: http
//     Host: localhost:8080
//     BasePath:
//     Version: v1.0.0
//     License: None
//     Contact: Nafiul-Quddus<quddusjunior1916@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - base64
//
//     SecurityDefinitions:
//     base64:
//          type: apiKey
//          name: ar5go-app-key
//          in: header
// swagger:meta
package openapi

import (
	"image"
	"visitor-management-system/types"
	//"time"
	"visitor-management-system/model"
)

// swagger:model ClientError
type ClientError struct {
	// Message of the error
	// in: string
	// example: bad_request
	Message string `json:"message"`
}

// swagger:model ServerError
type ServerError struct {
	// Message of the error
	// in: string
	// example: server_error
	Message string `json:"message"`
}

// swagger:model UnAuthorized
type UnAuthorized struct {
	// Message of the error
	// in: string
	// example: unAuthorized
	Message string `json:"message"`
}

// swagger:model Genericsuccess
type Genericsuccess struct {
	// Message of the success
	// in: string
	// example: successfull
	Message string `json:"message"`
}
type CompanyPayload struct {
	// in: string
	// example: vivasoft
	CompanyName string `json:"company_name" validate:"required,min=2,max=30"`
	// in: string
	// example: banani
	Address string `json:"address"`
	// in: string
	// example: nafiul quddus
	SubscriberName string `json:"subscriber_name"`
	// in: string
	// example: quddusjunior1916@gmail.com
	SubscriberEmail string              `json:"subscriber_email" validate:"required,email"`
	Subscription    SubscriptionPayload `gorm:"ForeignKey:CompanyId,constraint:OnUpdate:CASCADE"`
}

type SubscriptionPayload struct {
	// in: string
	// example: silver
	Subscription_type string `json:"subscription_type" validate:"required,eq=free|eq=silver|eq=premium"`
}

// Payload for create a Subscriber
// swagger:parameters CreateSub
type CreateSubPayload struct {
	// in : body
	Body CompanyPayload
}

// response after a Subscriber created
// swagger:response CreateSubResponse
type CreateSubResponse struct {
	// in: body
	Body model.Company
}
type UserD struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Type     string `json:"user_type" validate:"required,eq=Admin|eq=Official"`
}

// Payload for change a subscription
// swagger:parameters ChangeSub
type ChangeSubPayload struct {
	// in : body
	Body SubscriptionPayload
}

// Payload for user login
// swagger:parameters LoginDetails
type LoginDetails struct {
	// in : body
	Body UserD
}

// response after a user login
// swagger:response LoginSuccess
type LoginSuccessResponse struct {
	// in: body
	Body types.Token
}

// Payload for changing user password
// swagger:parameters ChangePassword
type ChangePassword struct {
	//in:body
	Body types.Password
}

type UserCreate struct {
	Name  string `json:"name" validate:"required,min=2,max=30"`
	Email string `json:"email" validate:"required,email"`
}

// Payload for changing user password
// swagger:parameters CreateUser
type CreateUser struct {
	//in:body
	Body UserCreate
}

// Payload for deleting a user
// swagger:parameters DeleteUser
type DeleteUser struct {
	//in:body
	Body ID
}

type ID struct {
	Id int `json:"id"`
}

// response about all user details
// swagger:response UserDetails
type UserDetails struct {
	// in: body
	Body []model.User
}

type RegVis struct {
	Name                  string      `json:"name" validate:"required,min=2,max=30"`
	Email                 string      `json:"email" validate:"required,email"`
	Phone                 string      `json:"phone" validate:"required,number"`
	Address               string      `json:"address"`
	Image                 image.Image `json:"image"`
	CompanyRepresentating string      `jsons:"company_rep"`
}

// Payload for creating a visitor
// swagger:parameters CreateVisitor
type CreateVisitor struct {
	//in:body
	Body RegVis
}
type Phone struct {
	Phone string `json:"phone"`
}
type Id struct {
	Id int `json:"v_id"`
}

// Payload for searching a visitor
// swagger:parameters Isregistered
type Isregistered struct {
	//in:body
	Body Phone
}
type VisitorsearchDetails struct {
	Id        int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name      string `json:"name" validate:"required,min=2,max=30"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required,number"`
	Address   string `json:"address"`
	CompanyId int
	ImageName string

	ImagePath             string
	CompanyRepresentating string `json:"company_rep"`
}

// response for searching a visitor
// swagger:response Visitordetails
type Visitordetails struct {
	//in:body
	Body VisitorsearchDetails
}

// track visitor only contains todays data if the  visitor hava multiple visits
// swagger:response LogResponse
type LogResponse struct {
	//in:body
	Body []model.Visitor
}

type PayloadCheckin struct {
	Image        image.Image `json:"image"`
	VId          int         `json:"v_id" validate:"required,number"`
	Status       string      `json:"status" validate:"required,eq=Arrived|eq=WillArrive|eq=Left"`
	Purpose      string      `json:"purpose" validate:"required,min=7,max=40"`
	AppointedTo  string      `json:"appointed_to" validate:"required,min=7,max=40"`
	FloorNumber  int         `json:"floor_number" validate:"required,number"`
	LuggageToken string      `json:"luggage_token"`
}

// Payload for searching a visitor
// swagger:parameters Checkin
type Checkin struct {
	//in:body
	Body PayloadCheckin
}

// all visitors
// swagger:response AllVisitor
type AllVisitor struct {
	//in:body
	Body []VisitorsearchDetails
}

// Payload for searching a visitor details
// swagger:parameters AllDetails
type AllDetails struct {
	//in:body
	Body ID
}

// all visits
// swagger:response Alltrackdetaails
type Alltrackdetaails struct {
	//in:body
	Body []model.TrackVisitor
}
