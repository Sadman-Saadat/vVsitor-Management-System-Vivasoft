package repository

import (
	//"fmt"
	"visitor-management-system/model"
	"visitor-management-system/types"
)

func CreateUser(user *model.User) error {
	err := db.Create(&user).Error
	return err
}

func GetAllUsers(id int) ([]types.UserDetails, error) {
	join_sql := "SELECT users.id,users.name,users.email,users.sub_domain ,users.company_id,users.branch_id,users.user_type,branches.branch_name,branches.address FROM users LEFT JOIN branches ON users.branch_id = branches.id WHERE users.company_id = ?"
	var official_users []types.UserDetails
	//err := db.Where("company_id = ?", id).Find(&official_users).Error
	err := db.Raw(join_sql, id).Scan(&official_users).Error
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

func GetBranchDetails(id int, bid int) (*model.Branch, error) {

	var branch model.Branch
	err := db.Model(&branch).Where("company_id = ? AND id = ?", id, bid).Find(&branch).Error
	return &branch, err
}
