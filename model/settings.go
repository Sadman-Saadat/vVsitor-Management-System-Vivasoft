package model

type Setting struct {
	Id        int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CompanyId int    `json:"company_id"`
	Email     string `json:"email" validate:"required,eq=yes|eq=no"`
	Image     string `json:"image" validate:"required,eq=yes|eq=no"`
}
