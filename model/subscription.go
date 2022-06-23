package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Subscription struct {
	Id                 int    `gorm:"primary_key;AUTO_INCREMENT"`
	CompanyId          int    `json:"company_id"`
	Subscription_type  string `json:"subscription_type" validate:"required,eq=free|eq=silver|eq=premium"`
	Subscription_start time.Time
	Subscription_end   time.Time
}
