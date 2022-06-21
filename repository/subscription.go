package repository

import (
	//"fmt"
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
	err = db.Preload("OfficialUser").Find(&all_sub).Error
	return
}

func GetSubscriberByEmail(email string) (*model.Subscriber, error) {
	var admin model.Subscriber

	err := db.Where("email = ?", email).Find(&admin).Error
	return &admin, err
}

func UpdateSubscriber(subscriber *model.Subscriber) error {
	err := db.Save(&subscriber).Error
	return err
}
