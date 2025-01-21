package auth

import (
	"errors"
)

func ValidateOAuthToken(token string) (bool, error) {
	if token == "" {
		return false, errors.New("invalid token")
	}
	return true, nil
}
