package repository

import (
	//"fmt"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//"visitor-management-system/database"
	//"time"
	"visitor-management-system/model"
)

func CreateRecord(record *model.Record) error {
	err := db.Create(&record).Error
	return err
}
