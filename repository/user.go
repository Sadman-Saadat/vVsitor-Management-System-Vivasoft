package repository

import (
	//"fmt"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//"visitor-management-system/database"
	"visitor-management-system/model"
)

func CreateUser(user *model.User) error {
	err := db.Create(&user).Error
	return err
}

func GetAllOfficialUsers(id int) ([]model.User, error) {
	var official_users []model.User
	err := db.Where("subscriber_id = ?", id).Find(&official_users).Error
	return official_users, err
}

func DeleteOfficialUser(user *model.User) error {
	err := db.Delete(user).Error
	return err
}

func UpdateOfficialUser(user *model.User) error {
	err := db.Save(&user).Error
	return err
}

func GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	err := db.Where("email = ?", email).Find(&user).Error
	return &user, err
}
