package views

import "time"

// register
type Register struct {
	Id        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

type Login struct {
	Token string `json:"token"`
}

type UpdateUser struct {
	Message string `json:"message"`
}
