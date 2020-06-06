package errors

import "errors"

// nolint
var (
	UnableToCreateUser = errors.New("unable to create user")
	UserAlreadyExists  = errors.New("user already exist")
)
