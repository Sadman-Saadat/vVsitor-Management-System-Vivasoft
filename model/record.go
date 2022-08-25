package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Record struct {
	Id                    int `gorm:"primary_key;AUTO_INCREMENT"`
	VId                   int
	Name                  string
	Email                 string
	Phone                 string
	Address               string
	ImagePath             string
	AppointedTo           string
	AppointedToPhone      string
	LuggageToken          string
	CompanyRepresentating string
	Status                string
	CompanyId             int
	BranchId              int
	BranchName            string
	Date                  time.Time
	CheckIn               string
	CheckOut              string
}
