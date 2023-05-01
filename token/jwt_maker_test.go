package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/ttanik/simplebank/util"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	assert.NoError(t, err)
	username := util.RandomOwner()
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	assert.NoError(t, err)
	assert.NotEmpty(t, payload)

	assert.NotZero(t, payload.ID)
	assert.Equal(t, username, payload.Username)
	assert.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	assert.NoError(t, err)
	username := util.RandomOwner()

	token, err := maker.CreateToken(username, -time.Minute)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrExpiredToken.Error())
	assert.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	payload, err := NewPayload(util.RandomOwner(), time.Minute)
	assert.NoError(t, err)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	assert.NoError(t, err)
	maker, err := NewJWTMaker(util.RandomString(32))
	assert.NoError(t, err)
	payload, err = maker.VerifyToken(token)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrInvalidToken.Error())
	assert.Nil(t, payload)

}

