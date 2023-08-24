package util

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("Failed to hash password")
	}

	return string(hashed_password), nil
}

func CheckPassword(password string, hashed_password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(password))
}
