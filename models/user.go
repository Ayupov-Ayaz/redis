package models

import "errors"

var (
	UserNotFound = errors.New("User not found")
	AlreadyExist = errors.New("user with email already exist")
)

type User struct {
	ID       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name" validate:"required"`
	Email    string `json:"email" db:"email" validate:"required,email"`
	Password string `json:"password" db:"password" validate:"required,min:6"`
}
