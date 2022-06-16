package types

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Type     string `json:"user_type" validate:"required,eq=Admin|eq=Official"`
}
