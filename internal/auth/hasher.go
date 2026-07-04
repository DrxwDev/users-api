package auth

import (
	"runtime"

	"github.com/alexedwards/argon2id"

	errormsg "github.com/DrxwDev/users-api/internal/errorMsg"
)

type HashService interface {
	Hash(password string) (string, error)
	VerifyHash(password, hash string) error
}

type argonHasher struct {
	params *argon2id.Params
}

func NewArgonHasher() HashService {
	return &argonHasher{
		params: &argon2id.Params{
			Memory:      64 * 1024,
			Iterations:  2,
			Parallelism: uint8(runtime.NumCPU()),
			SaltLength:  16,
			KeyLength:   32,
		},
	}
}

func (auth argonHasher) Hash(password string) (string, error) {
	if password == "" {
		return "", errormsg.ErrPasswordNotProvide
	}

	return argon2id.CreateHash(password, auth.params)
}

func (auth argonHasher) VerifyHash(password, hash string) error {
	if password == "" {
		return errormsg.ErrPasswordNotProvide
	}

	if hash == "" {
		return errormsg.ErrHashNotProvided
	}

	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return errormsg.ErrVerifyingPasswordAndHash
	}

	if !match {
		return errormsg.ErrPasswordNotMatch
	}

	return nil
}
