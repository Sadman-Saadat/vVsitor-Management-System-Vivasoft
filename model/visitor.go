package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"image"
)

type Visitor struct {
	Id                    int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name                  string `json:"name" validate:"required,min=2,max=30"`
	Email                 string `json:"email" validate:"required,email"`
	Phone                 string `json:"phone" validate:"required,number"`
	Address               string `json:"address"`
	CompanyId             int
	Status                string `json:"status" validate:"required,eq=Arrived|eq=WillArrive|eq=Left"`
	ImageName             string
	ImagePath             string
	CompanyRepresentating string         `json:"company_rep"`
	TrackVisitors         []TrackVisitor `gorm:"ForeignKey:VId"`
}

type TrackVisitor struct {
	Id           int `gorm:"primary_key;AUTO_INCREMENT"`
	VId          int `json:"v_id" validate:"required,number"`
	CompanyId    int
	Purpose      string `json:"purpose" validate:"required,min=7,max=40"`
	AppointedTo  string `json:"appointed_to" validate:"required,min=7,max=40"`
	FloorNumber  int    `json:"floor_number" validate:"required,number"`
	LuggageToken string `json:"luggage_token"`
	ImagePath    string
	Date         string
	CheckIn      string
	CheckOut     string
}
