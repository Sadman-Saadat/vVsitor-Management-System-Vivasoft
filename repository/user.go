package repository

import (
	"visitor-management-system/model"
)

func CreateUser(user *model.User) error {
	err := db.Create(&user).Error
	return err
}

func GetAllUsers(id int) ([]model.User, error) {
	var official_users []model.User
	err := db.Where("company_id = ?", id).Find(&official_users).Error
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

func GetUserByEmail(email string, id int) (*model.User, error) {
	var user model.User

	err := db.Where("company_id = ? AND email = ?", id, email).Find(&user).Error
	return &user, err
}
