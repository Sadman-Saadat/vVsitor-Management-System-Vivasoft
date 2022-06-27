package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
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
	Token                 string
	CompanyRepresentating string
	Date                  time.Time
	CheckIn               time.Time
	checkOut              time.Time
}
