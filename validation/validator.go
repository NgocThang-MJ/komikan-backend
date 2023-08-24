package validation

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func validateLen(key string, value string, minLen int, maxLen int) error {
	n := len(value)
	if n < minLen || n > maxLen {
		return fmt.Errorf("%s must contain from %d-%d characters", key, minLen, maxLen)
	}
	return nil
}

func validateFullName(value string) error {
	if err := validateLen("Full name", value, 3, 100); err != nil {
		return err
	}

	if !isValidFullName(value) {
		return fmt.Errorf("Must contain only letters or spaces")
	}
	return nil
}

func validateUsername(value string) error {
	if err := validateLen("Username", value, 3, 100); err != nil {
		return err
	}

	if !isValidUsername(value) {
		return fmt.Errorf("Username must contain only letters, digits and underscore")
	}
	return nil
}

func validatePassword(value string) error {
	return validateLen("Password", value, 6, 100)
}

func validateEmail(value string) error {
	if err := validateLen("Email", value, 3, 100); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("Not a valid email address")
	}

	return nil
}
