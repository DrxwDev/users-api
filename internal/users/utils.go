package users

import (
	"strings"

	"github.com/google/uuid"

	errormsg "github.com/DrxwDev/users-api/internal/errorMsg"
)

func validateEmail(email string) error {
	if email == "" {
		return errormsg.ErrEmailNotProvided
	}

	if len(email) > 100 {
		return errormsg.ErrEmailTooLong
	}

	if !strings.Contains(email, "@") {
		return errormsg.ErrInvalidEmailFormat
	}

	return nil
}

func validatePassword(password string) error {
	if password == "" {
		return errormsg.ErrPasswordNotProvide
	}

	if len(password) > 250 {
		return errormsg.ErrPasswordTooLong
	}

	if len(password) < 8 {
		return errormsg.ErrPasswordTooShort
	}

	return nil
}

func validateUserID(id string) error {
	if id == "" {
		return errormsg.ErrUserIDNotProvided
	}

	return nil
}

func parseUserID(id string) (uuid.UUID, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, errormsg.ErrUnableToParseUserID
	}

	return uid, nil
}
