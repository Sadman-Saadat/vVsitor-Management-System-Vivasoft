package repository

import (
	//"visitor-management-system/database"
	"visitor-management-system/model"
)

func CreateMasterAdmin(master *model.MasterAdmin) (*model.MasterAdmin, error) {
	err := db.Create(&master).Error
	return master, err
}

func GetMasterAdminByEmail(email string) (*model.MasterAdmin, error) {
	var master *model.MasterAdmin
	err := db.Where("email=?", email).Find(&master).Error
	return master, err
}

func CreatePackage(packges *model.Package) (*model.Package, error) {
	err := db.Create(&packges).Error
	return packges, err
}
