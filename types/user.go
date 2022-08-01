package types

type User struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
	BranchName string `json:"branch_name"`
	Name       string `json:"name"`
	BranchId   int    `json:"branch_id"`
}

type Password struct {
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type Token struct {
	User_Token        string
	User_Refreshtoken string
}

type UserDetails struct {
	Id         int    `gorm:"primary_key;AUTO_INCREMENT"`
	Name       string `json:"name" validate:"required,min=2,max=30"`
	Email      string `json:"email" validate:"required,email"`
	SubDomain  string
	CompanyId  int
	BranchId   int
	UserType   string `json:"user_type" validate:"eq=Admin|eq=Official"`
	BranchName string `json:"branch_name" validate:"required"`
	Address    string `json:"address"`
}
