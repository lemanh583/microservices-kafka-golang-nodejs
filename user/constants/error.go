package constants

import "errors"

var (
	ERROR_NOT_FOUND = errors.New("not found record")
	ERROR_EXIST     = errors.New("exists")
	ERROR_SIGN_IN   = errors.New("wrong email or password")
)
