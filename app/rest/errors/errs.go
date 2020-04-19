package errors

import "errors"

var (
	PasswordNotMatch = errors.New("Invalid password. ")
	InvalidToken     = errors.New("Invalid token. ")
)
