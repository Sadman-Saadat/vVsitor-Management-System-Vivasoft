package repository

import (
	//"visitor-management-system/database"
	"visitor-management-system/model"
)

func CreateNewSettings(settings *model.Setting) (*model.Setting, error) {
	err := db.Create(&settings).Error
	return settings, err
}

func Setting(id int) (model.Setting, error) {
	var settings model.Setting
	err := db.Where("company_id = ?", id).Find(&settings).Error
	return settings, err
}

func UpdateSettings(settings *model.Setting) error {
	err := db.Where("company_id = ?", settings.CompanyId).Save(&settings).Error
	return err
}
