package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hash), nil
}

func CheckPassword(pass string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
}
