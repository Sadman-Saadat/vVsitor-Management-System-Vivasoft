package model

type UserBranchRelation struct {
	Id        int `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CompanyId int `json:"company_id"`
	UserId    int `json:"user_id" validate:"required"`
	BranchId  int `json:"branch_id" validate:"required"`
}
