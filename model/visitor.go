package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Visitor struct {
	Id          int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name        string `json:"name" validate:"required,min=2,max=30"`
	Email       string `json:"email" validate:"required,email"`
	Phone       string `json:"phone" validate:"required,number"`
	Purpose     string `json:"purpose" validate:"required,min=7,max=40"`
	MeetingWith string `json:"meeting_with" validate:"required,email"`
	CheckIn     string
	CheckOut    string
}
