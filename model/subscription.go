package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Subscription struct {
	CompanyId          int    `json:"company_id" gorm:"primary_key"`
	Subscription_type  string `json:"subscription_type" validate:"required,eq=free|eq=silver|eq=premium|eq=cancel"`
	Subscription_start time.Time
	Subscription_end   time.Time
}
