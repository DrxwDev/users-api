package errormsg

import (
	"errors"
)

var (
	ErrUnableToSaveUser    = errors.New("unable to save user")
	ErrUserNotFound        = errors.New("user not found")
	ErrEmailNotProvided    = errors.New("user email not provided")
	ErrInvalidEmailFormat  = errors.New("invalid email format")
	ErrEmailTooLong        = errors.New("email is too long 100 max charaters")
	ErrPasswordNotProvided = errors.New("password not provided")
	ErrPasswordTooLong     = errors.New("password too long")
	ErrPasswordTooShort    = errors.New("password too short")
	ErrUserIDNotProvided   = errors.New("user id not provided")
	ErrUnableToParseUserID = errors.New("unable to parse user id")
	ErrUserAlreadyExists   = errors.New("unable to create user, user already exists")
)
