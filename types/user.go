package types

type User struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required"`
	BranchName string `json:"branch_name"`
	Name       string `json:"name"`
}

type Password struct {
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type Token struct {
	User_Token        string
	User_Refreshtoken string
}
