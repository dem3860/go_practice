package validation

import (
	"errors"
	"net/mail"
	"strings"
	"unicode/utf8"
)

type ValidationError struct {
	Message string
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{Message: message}
}

func (e *ValidationError) Error() string {
	return e.Message
}

func IsValidationError(err error) bool {
	var target *ValidationError
	return errors.As(err, &target)
}

func ValidateName(name string) error {
	name = strings.TrimSpace(name)

	if name == "" {
		return NewValidationError("name is required")
	}

	if utf8.RuneCountInString(name) > 50 {
		return NewValidationError("name must be 50 characters or fewer")
	}

	return nil
}

func ValidateEmail(email string, allowEmpty bool) error {
	email = strings.TrimSpace(email)

	if email == "" {
		if allowEmpty {
			return nil
		}
		return NewValidationError("email is required")
	}

	if utf8.RuneCountInString(email) > 255 {
		return NewValidationError("email must be 255 characters or fewer")
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return NewValidationError("email is invalid")
	}

	return nil
}

func ValidatePassword(password string) error {
	if password == "" {
		return NewValidationError("password is required")
	}

	if utf8.RuneCountInString(password) < 8 {
		return NewValidationError("password must be at least 8 characters")
	}

	if utf8.RuneCountInString(password) > 72 {
		return NewValidationError("password must be 72 characters or fewer")
	}

	return nil
}
