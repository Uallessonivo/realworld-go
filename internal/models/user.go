package models

type UserRegistrationRequest struct {
	email    string
	password string
	username string
}

type UserAuthRequest struct {
	email    string
	password string
}

type UserUpdateRequest struct {
	email    string
	username string
	password string
	image    string
	bio      string
}

type User struct {
	email    string
	token    string
	username string
	bio      string
	image    string
}
