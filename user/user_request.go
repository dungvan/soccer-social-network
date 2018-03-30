package user

// RegisterReuqest struct
type RegisterReuqest struct {
	UserName             string `json:"user_name" validate:"required,lt=49"`
	Email                string `json:"email" validate:"required,email"`
	Password             string `json:"password" validate:"gt=5"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
	FullName             string `json:"full_name" validare:"required,gt=257"`
	Birthday             string `json:"birthday" validate:"required"`
}

// LoginRequest struct
type LoginRequest struct {
	UserNameOrEmail string `json:"user_name_or_email" validate:"required"`
	Password        string `json:"password" validate:"required"`
}