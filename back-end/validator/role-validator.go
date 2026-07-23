package validator

import (
	"errors"
)

func ValidateRole(role string) error {
	switch role {
	case "admin":
		return nil

	case "user":
		return nil

	default:
		return errors.New("invalid role")
	}
}