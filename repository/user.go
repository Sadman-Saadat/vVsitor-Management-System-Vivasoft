package repository

import (
	"fmt"
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

func GetUserByEmail(email string, subdomain string) (*model.User, error) {
	var user model.User

	err := db.Where("sub_domain = ? AND email = ?", subdomain, email).Find(&user).Error
	return &user, err
}

func GetBranchDetails(id int, name string) (*model.Branch, error) {
	fmt.Println(id)
	fmt.Println(name)
	var branch model.Branch
	err := db.Model(&branch).Where("company_id = ? AND branch_name = ?", id, name).Find(&branch).Error
	return &branch, err
}
