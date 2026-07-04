// Package errormsg
package errormsg

import (
	"errors"
)

var (
	ErrPasswordNotProvide       = errors.New("password not provided by the user")
	ErrHashNotProvided          = errors.New("hash not provided by the user")
	ErrVerifyingPasswordAndHash = errors.New("unable to compare password and hash")
	ErrPasswordNotMatch         = errors.New("invalid password")
)
