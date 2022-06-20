package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"image"
)

type Visitor struct {
	Id                    int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name                  string `json:"name" validate:"required,min=2,max=30"`
	Email                 string `json:"email" validate:"required,email"`
	Phone                 string `json:"phone" validate:"required,number"`
	Address               string `json:"address"`
	ImageName             string
	ImagePath             string
	CompanyRepresentating string         `json:"company_rep"`
	TrackVisitors         []TrackVisitor `gorm:"ForeignKey:VId"`
}

type TrackVisitor struct {
	Id       int `gorm:"primary_key;AUTO_INCREMENT"`
	VId      int
	Email    string `json:"email"`
	Purpose  string `json:"purpose" validate:"required,min=7,max=40"`
	Token    string `json:"token"`
	CheckIn  string
	CheckOut string
}
