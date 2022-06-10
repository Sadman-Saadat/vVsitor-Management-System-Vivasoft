package repository

import (
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"visitor-management-system/database"
	"visitor-management-system/model"
)

var db = database.GetDB()

func CreateSub(subscriber *model.Subscription) error {

	err := db.Create(&subscriber).Error
	return err

}
