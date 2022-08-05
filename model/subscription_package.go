package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"time"
)

type Package struct {
	Id                int    `gorm:"primary_key;AUTO_INCREMENT"`
	Subscription_type string `json:"subscription_type" validate:"required,eq=free|eq=silver|eq=premium|eq=cancel"`
	Duration          int
}
