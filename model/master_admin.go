package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MasterAdmin struct {
	Id       int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name     string `json:"name" validate:"required,min=2,max=30"`
	Email    string `json:"email" validate:"required,email"`
	UserType string `json:"user_type" validate:"required"`
	Password string
}
