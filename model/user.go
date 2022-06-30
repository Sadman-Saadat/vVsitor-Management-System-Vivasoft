package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id        int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string `json:"name" validate:"required,min=2,max=30"`
	Email     string `json:"email" validate:"required,email"`
	CompanyId int
	UserType  string `json:"user_type" validate:"required,eq=Admin|eq=Official"`
	Password  string
}
