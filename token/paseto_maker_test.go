package token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/ttanik/simplebank/util"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
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

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
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
