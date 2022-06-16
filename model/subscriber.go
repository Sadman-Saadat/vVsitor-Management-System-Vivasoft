package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Subscriber struct {
	Id                 int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name               string `json:"name" validate:"required,min=2,max=30"`
	Email              string `json:"email" validate:"required,email"`
	Address            string `json:"address"`
	Subscription_type  string `json:"subscription_type" validate:"required,eq=general|eq=silver|eq=premium"`
	Password           string
	Subscription_start string
	Subscription_end   string
	Token              string
	RefreshToken       string
}
