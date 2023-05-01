package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)
	hashed, err := HashPassword(password)
	assert.NoError(t, err)

	err = CheckPassword(password, hashed)
	assert.NoError(t, err)

	wrongPass := "cu"
	err = CheckPassword(wrongPass, hashed)
	assert.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
