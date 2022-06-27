package repository

import (
	//"fmt"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"visitor-management-system/database"
	"visitor-management-system/model"
)

var db = database.GetDB()

func RegisterCompany(company *model.Company) error {
	err := db.Create(&company).Error
	return err
}

func GetAllSubscriber() (all_company []model.Company, err error) {
	err = db.Preload("User").Find(&all_company).Error
	return
}

// func GetSubscriberByEmail(email string) (*model.Company, error) {
// 	var admin model.Company

// 	err := db.Where("email = ?", email).Find(&admin).Error
// 	return &admin, err
// }

func UpdateUser(user *model.User) error {
	err := db.Save(&user).Error
	return err
}

func ChangeSubscription(sub *model.Subscription) error {
	// err := db.Where("company_id = ?", sub.CompanyId).Delete(&sub).Error
	// err = db.Save(&sub).Error
	err := db.Model(&sub).Where("company_id = ?", sub.CompanyId).Update("subscription_type", sub.Subscription_type).Error
	return err
}

func CancelSubscription(sub *model.Subscription) error {
	err := db.Model(&sub).Where("company_id = ?", sub.CompanyId).Update("subscription_type", "free").Error
	return err
}
