package repository

import (
	//"fmt"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//"visitor-management-system/database"
	"visitor-management-system/model"
)

func CreateVisitor(visitor *model.Visitor) error {
	err := db.Create(&visitor).Error
	return err
}

func GetAllVisitor() (visitor []*model.Visitor, err error) {
	err = db.Preload("TrackVisitors").Find(&visitor).Error
	return
}

func GetVisitor(visitor *model.Visitor) (*model.Visitor, error) {
	err := db.Find(&visitor).Error
	return visitor, err
}

func GetVisitorDetails(visitor *model.Visitor) (*model.Visitor, error) {
	err := db.Preload("TrackVisitors").Find(&visitor).Error
	return visitor, err
}

func UpdateVisitor(visitor *model.Visitor) error {
	err := db.Save(&visitor).Error
	return err
}
