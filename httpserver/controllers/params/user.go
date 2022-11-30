package params

// register
type Register struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// login
type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// balance
type UpdateUser struct {
	Balance uint `json:"balance" validate:"required, min=0, max=100000000"`
}
