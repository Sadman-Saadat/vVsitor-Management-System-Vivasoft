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

func Search(visitor *model.Visitor) (*model.Visitor, error) {
	err := db.Where("phone = ?", visitor.Phone).Find(&visitor).Error
	return visitor, err
}

func CheckIn(info *model.TrackVisitor) error {
	err := db.Create(&info).Error
	return err
}

func CountPresentVisitor() (int, error) {
	var count int
	var visitor []*model.Visitor
	err := db.Where("status = ?", "Arrived").Or("status = ?", "WillArrive").Find(&visitor).Count(&count).Error
	return count, err
}
