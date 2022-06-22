package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Subscription struct {
	Id                 int    `gorm:"primary_key;AUTO_INCREMENT"`
	CompanyId          int    `json:"company_id"`
	Subscription_type  string `json:"subscription_type" validate:"required,eq=free|eq=silver|eq=premium"`
	Subscription_start string
	Subscription_end   string
}
