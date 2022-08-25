package repository

import (
	"visitor-management-system/database"
	"visitor-management-system/model"
	//"visitor-management-system/types"
)

var db = database.GetDB()

func RegisterCompany(company *model.Company) (*model.Company, error) {
	err := db.Create(&company).Error
	return company, err
}

func GetAllSubscriber() (all_company []model.Company, err error) {
	err = db.Preload("User").Find(&all_company).Error
	return
}

func UpdateUser(user *model.User) error {
	err := db.Save(&user).Error
	return err
}

func IsCompanyValid(name string, subdomain string) (int64, error) {
	var existing_company []*model.Company
	var count int64
	err := db.Where("company_name = ?", name).Or("sub_domain = ?", subdomain).Find(&existing_company).Count(&count).Error
	return count, err
}

func GetPackageById(id int) (*model.Package, error) {
	var pack *model.Package
	err := db.Model(&model.Package{}).Where("id=?", id).Find(&pack).Error
	return pack, err
}

func SetAdminPassword(id int, password string) (*model.User, error) {
	var user *model.User
	err := db.Model(&model.User{}).Where("id =?", id).Update("password", password).Find(&user).Error
	return user, err
}

func SetCompanyStatus(id int) error {
	err := db.Model(&model.Company{}).Where("id=?", id).Update("status", true).Error
	return err
}
