package model

type Role string

type UserDTO struct {
	Name     string `json:"name" db:"name"`
	Lastname string `json:"lastname" db:"lastname"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}
