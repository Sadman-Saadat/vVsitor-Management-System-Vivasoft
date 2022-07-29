package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Visitor struct {
	Id        int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name      string `json:"name" validate:"required,min=2,max=30"`
	Email     string `json:"email" validate:"email"`
	Phone     string `json:"phone" validate:"required,number"`
	Address   string `json:"address" validate:"required"`
	CompanyId int
	BranchId  int
	ImageName string

	ImagePath             string
	CompanyRepresentating string         `json:"company_rep"`
	TrackVisitors         []TrackVisitor `gorm:"ForeignKey:VId;references:Id"`
}

type TrackVisitor struct {
	Id               int `gorm:"primary_key;AUTO_INCREMENT"`
	VId              int `json:"v_id" validate:"required,number"`
	CompanyId        int
	BranchId         int
	Status           string `json:"status" validate:"required,eq=Arrived|eq=WillArrive|eq=Left"`
	Purpose          string `json:"purpose" validate:"required,min=7,max=40"`
	AppointedTo      string `json:"appointed_to" validate:"required,min=7,max=40"`
	AppointedToPhone string `json:"appointed_to_phone" validate:"number"`
	FloorNumber      int    `json:"floor_number" validate:"number"`
	LuggageToken     string `json:"luggage_token"`
	ImagePath        string
	Date             time.Time
	CheckIn          string
	CheckOut         string
}
