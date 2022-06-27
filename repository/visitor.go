package repository

import (
	//"fmt"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//"visitor-management-system/database"
	"time"
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

func CountPresentVisitor(id int) (int, error) {
	var count int
	today := time.Now().Local().Format("2006-01-02")
	val := "Arrived"
	var visitor []*model.TrackVisitor
	err := db.Where("status = ? AND date=? AND company_id = ?", val, today, id).Find(&visitor).Count(&count).Error
	return count, err
}
