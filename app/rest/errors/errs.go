package errors

import "errors"

var (
	PasswordNotMatch  = errors.New("Invalid password. ")
	InvalidToken      = errors.New("Invalid token. ")
	UserAlreadyExists = errors.New("A user with this mail already exists. ")
)
