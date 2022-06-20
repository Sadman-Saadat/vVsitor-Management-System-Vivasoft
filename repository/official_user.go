package repository

import (
	//"fmt"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//"visitor-management-system/database"
	"visitor-management-system/model"
)

func CreateOfficialUser(official_user *model.OfficialUser) error {
	err := db.Create(&official_user).Error
	return err
}

func GetAllOfficialUsers(id int) ([]model.OfficialUser, error) {
	var official_users []model.OfficialUser
	err := db.Where("subscriber_id = ?", id).Find(&official_users).Error
	return official_users, err
}

func DeleteOfficialUser(user *model.OfficialUser) error {
	err := db.Delete(user).Error
	return err
}

func UpdateOfficialUser(user *model.OfficialUser) error {
	err := db.Save(&user).Error
	return err
}

func GetOfficialUserByEmail(email string) (*model.OfficialUser, error) {
	var user model.OfficialUser

	err := db.Where("email = ?", email).Find(&user).Error
	return &user, err
}
