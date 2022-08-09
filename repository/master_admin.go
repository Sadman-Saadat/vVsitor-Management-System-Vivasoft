package repository

import (
	"fmt"
	"strconv"
	"time"
	"visitor-management-system/model"
	"visitor-management-system/types"
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

func GetCompanyList(limit int, offset int, search string, status string, package_id int) ([]*types.CompanyDetails, int64, error) {
	var company []*types.CompanyDetails
	var count int64
	var bool_search bool
	dbmodel := db.Model(&model.Company{}).Select("companies.*, packages.subscription_type").Joins("left join packages on packages.id = companies.package_id")
	if search != "" {
		search += fmt.Sprintf("%s", "%")
		dbmodel = dbmodel.Where("company_name LIKE ? OR sub_domain LIKE ? ", search, search)
	}
	if status != "" {
		bool_search, _ = strconv.ParseBool(status)
		dbmodel = dbmodel.Where("status = ?", bool_search)
	}
	if package_id != 0 {
		dbmodel = dbmodel.Where("package_id = ?", package_id)
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

func GetCompanyAdminlist(limit int, offset int, search string) ([]*types.AdminDetails, int64, error) {
	var admins []*types.AdminDetails
	var count int64
	dbmodel := db.Model(&model.User{}).Select("users.*,companies.company_name,companies.address").Where("user_type =?", "Admin").Joins("left join companies on companies.id = users.company_id")
	if search != "" {
		search += fmt.Sprintf("%s", "%")
		dbmodel = dbmodel.Where("companies.company_name LIKE ? OR users.name LIKE ? OR users.email LIKE ? OR users.sub_domain LIKE ?", search, search, search, search)
	}

	dbmodel = dbmodel.Count(&count)
	dbmodel = dbmodel.Limit(limit).Offset(offset).Find(&admins)
	err := dbmodel.Error
	return admins, count, err
}

func DeletePackage(id int) (error, error) {
	err1 := db.Model(&model.Package{}).Where("id = ?", id).Delete(&model.Package{}).Error
	err2 := db.Model(&model.PackageFeatures{}).Where("package_id = ?", id).Delete(&model.PackageFeatures{}).Error
	return err1, err2

}

func UpdateFeatures(new_features *model.PackageFeatures) error {
	err := db.Save(&new_features).Error
	return err
}

func UpdateCompanyStatus(company_id int, status bool) error {
	err := db.Model(&model.Company{}).Where("id =?", company_id).Update("status", status).Error
	return err
}

func ChangeSubscription(sub *types.ChangeSubscription, start time.Time, end time.Time) error {

	err := db.Model(&model.Company{}).Where("id = ?", sub.CompanyId).Updates(map[string]interface{}{"package_id": sub.PackageId, "subscription_start": start, "status": true, "subscription_end": end}).Error
	return err
}
