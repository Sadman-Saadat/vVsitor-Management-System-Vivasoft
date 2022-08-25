package repository

import (
	//"visitor-management-system/database"
	"visitor-management-system/model"
)

func CreateNewBranch(branch *model.Branch) (*model.Branch, error) {
	err := db.Create(&branch).Error
	return branch, err
}

func BranchList(id int) (*[]model.Branch, error) {
	var all_branch *[]model.Branch
	err := db.Where("company_id = ?", id).Find(&all_branch).Error
	return all_branch, err
}

func UpdateBranch(branch *model.Branch) error {
	err := db.Save(&branch).Error
	return err
}

func DeleteBranch(company_id int, id int) error {
	var branch model.Branch
	branch.Id = id
	err := db.Delete(&branch).Where("company_id = ?", company_id).Error
	return err
}

func IsBranchValid(company_id int, branch_name string) (int64, error) {
	var count int64
	var branch []*model.Branch
	err := db.Where("company_id = ? AND branch_name = ?", company_id, branch_name).Find(&branch).Count(&count).Error
	return count, err
}
