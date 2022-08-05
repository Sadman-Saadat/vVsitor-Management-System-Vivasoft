package repository

import (
	//"visitor-management-system/database"
	"fmt"
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

func GetCompanyList(limit int, offset int, search string) ([]*model.Company, int64, error) {
	var company []*model.Company
	var count int64
	dbmodel := db.Model(&model.Company{})
	if search != "" {
		search += fmt.Sprintf("%s", "%")
		dbmodel = dbmodel.Where("company_name LIKE ? AND sub_domain LIKE ?", search, search)
	}
	dbmodel = dbmodel.Count(&count)
	dbmodel = dbmodel.Limit(limit).Offset(offset).Find(&company)
	err := dbmodel.Error
	return company, count, err

}

func GetPackageList() ([]*model.Package, error) {
	var packages []*model.Package
	err := db.Model(&model.Package{}).Find(&packages).Error
	return packages, err
}
