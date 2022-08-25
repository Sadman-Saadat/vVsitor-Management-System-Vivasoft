package repository

import (
	"fmt"
	//"visitor-management-system/database"
	"visitor-management-system/model"
	//"visitor-management-system/types"
)

func CreateNewUserBranchRelation(relation *model.UserBranchRelation) error {
	fmt.Println(relation)
	err := db.Create(&relation).Error
	fmt.Println(relation)
	return err
}

func GetBranchRelation(user_id int, company_id int) ([]int64, error) {
	//var branch_ids *types.BranchIds
	var ids []int64
	sql := fmt.Sprintf("SELECT user_branch_relations.branch_id FROM user_branch_relations WHERE user_branch_relations.company_id = %d AND user_branch_relations.user_id = %d", company_id, user_id)
	err := db.Raw(sql).Find(&ids).Error
	//fmt.Println(ids)
	return ids, err
}

func GetBranchList(branch_ids []int64, company_id int) (*[]model.Branch, error) {
	var branch = new([]model.Branch)
	sql := fmt.Sprintf("SELECT * FROM branches WHERE branches.id IN ? AND branches.company_id = %d", company_id)
	err := db.Raw(sql, branch_ids).Find(&branch).Error
	return branch, err
}
