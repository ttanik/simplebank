package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ttanik/simplebank/util"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	assert.Nil(t, err)
	assert.Equal(t, arg.Balance, account.Balance)
	assert.Equal(t, arg.Owner, account.Owner)
	assert.Equal(t, arg.Currency, account.Currency)
	return account
}
func TestCreateAccout(t *testing.T) {
	_ = createRandomAccount(t)

}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	getAccount, err := testQueries.GetAccount(context.Background(), account1.ID)
	assert.NoError(t, err)
	assert.Equal(t, account1.ID, getAccount.ID)
	assert.Equal(t, account1.Currency, getAccount.Currency)
	assert.Equal(t, account1.Owner, getAccount.Owner)
	assert.Equal(t, account1.Balance, getAccount.Balance)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}
	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	assert.NoError(t, err)
	assert.Equal(t, account1.ID, account2.ID)
	assert.Equal(t, account1.Currency, account2.Currency)
	assert.Equal(t, account1.Owner, account2.Owner)
	assert.Equal(t, arg.Balance, account2.Balance)

}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	assert.NoError(t, err)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	assert.Error(t, err)
	assert.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	arg := ListAccountParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccount(context.Background(), arg)
	assert.NoError(t, err)
	for _, account := range accounts {
		assert.NotEmpty(t, account)
	}
}
