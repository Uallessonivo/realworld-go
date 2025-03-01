package models

import (
	"gorm.io/gorm"
)

type UserRegistrationRequest struct {
	Email    string
	Password string
	Username string
}

type UserAuthRequest struct {
	Email    string
	Password string
}

type UserUpdateRequest struct {
	Email    string
	Username string
	Password string
	Image    string
	Bio      string
}

type UserResponse struct {
	Email    string
	Token    string
	Username string
	Bio      string
	Image    string
}

type User struct {
	gorm.Model
	Username     string
	Email        string
	Bio          string
	Image        string
	PasswordHash string
}
