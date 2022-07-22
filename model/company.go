package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Company struct {
	Id              int          `gorm:"primary_key;AUTO_INCREMENT"`
	CompanyName     string       `json:"company_name" validate:"required,min=2,max=30"`
	Address         string       `json:"address"`
	SubscriberName  string       `json:"subscriber_name"`
	SubscriberEmail string       `json:"subscriber_email" validate:"required,email"`
	Subscription    Subscription `gorm:"foreignKey:CompanyId;references:Id"`
	User            []User       `gorm:"ForeignKey:CompanyId;references:Id"`
}
