package repository

import (
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"visitor-management-system/database"
	"visitor-management-system/model"
)

var db = database.GetDB()

func CreateSub(subscriber *model.Subscriber) error {
	err := db.Create(&subscriber).Error
	return err
}

func GetAllSubscriber() (all_sub []model.Subscriber, err error) {
	err = db.Find(&all_sub).Error
	return

}
