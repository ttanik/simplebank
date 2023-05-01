package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ttanik/simplebank/util"
)

func createRandomUser(t *testing.T) User {
	hashedPass, err := util.HashPassword(util.RandomString(6))
	assert.NoError(t, err)
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPass,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	assert.Nil(t, err)
	assert.Equal(t, arg.Username, user.Username)
	assert.Equal(t, arg.HashedPassword, user.HashedPassword)
	assert.Equal(t, arg.FullName, user.FullName)
	assert.Equal(t, arg.Email, user.Email)
	return user
}
func TestCreateUser(t *testing.T) {
	_ = createRandomUser(t)

}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	getuser, err := testQueries.GetUser(context.Background(), user1.Username)
	assert.NoError(t, err)
	assert.Equal(t, user1.Username, getuser.Username)
	assert.Equal(t, user1.HashedPassword, getuser.HashedPassword)
	assert.Equal(t, user1.FullName, getuser.FullName)
	assert.Equal(t, user1.Email, getuser.Email)
}
