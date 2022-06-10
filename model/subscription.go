package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Subscriber struct {
	gorm.Model
	Id                 int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name               string `form:"name"`
	Email              string `form:"email"`
	Address            string `form:"address"`
	Subscription_type  string `form:"subscription_type"`
	Subscription_start time.Time
	Subscription_end   time.Time
}
