package model

type Branch struct {
	Id         int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CompanyId  int    `json:"company_id"`
	BranchName string `json:"branch_name" validate:"required"`
	Address    string `json:"address" validate:"required"`
}
