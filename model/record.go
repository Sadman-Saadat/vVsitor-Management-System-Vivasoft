package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"time"
	//"image"
)

type Record struct {
	Id                    int `gorm:"primary_key;AUTO_INCREMENT"`
	VId                   int
	VisitorName           string
	VisitorEmail          string
	VisitorPhone          string
	VisitorAddress        string
	VisitorImagePath      string
	AppointedTo           string
	LuggageToken          string
	CompanyRepresentating string
	CompanyId             int
	Date                  string
	CheckIn               string
	CheckOut              string
}
