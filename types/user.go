package types

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Type     string `json:"user_type" validate:"required,eq=Admin|eq=Official"`
}

type Password struct {
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type Token struct {
	User_Token        string
	User_Refreshtoken string
}
