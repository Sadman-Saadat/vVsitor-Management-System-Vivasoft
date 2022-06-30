package repository

import (
	"visitor-management-system/model"
)

func CreateRecord(record *model.Record) error {
	err := db.Create(&record).Error
	return err
}
