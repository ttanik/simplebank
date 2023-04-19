package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ttanik/simplebank/util"
)

func createBaseAccounts(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	arg2 := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	account2, err2 := testQueries.CreateAccount(context.Background(), arg2)
	assert.Nil(t, err)
	assert.Equal(t, arg.Balance, account.Balance)
	assert.Equal(t, arg.Owner, account.Owner)
	assert.Equal(t, arg.Currency, account.Currency)
	assert.Nil(t, err2)
	assert.Equal(t, arg2.Balance, account2.Balance)
	assert.Equal(t, arg2.Owner, account2.Owner)
	assert.Equal(t, arg2.Currency, account2.Currency)
}

func createRandomTransfer(t *testing.T) Transfer {
	createBaseAccounts(t)
	arg := CreateTransferParams{
		FromAccountID: 1,
		ToAccountID:   2,
		Amount:        util.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	assert.Nil(t, err)
	assert.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	assert.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	assert.Equal(t, arg.Amount, transfer.Amount)
	return transfer
}
func TestCreateTransfer(t *testing.T) {
	_ = createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	getTransfer, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	assert.NoError(t, err)
	assert.Equal(t, transfer1.ID, getTransfer.ID)
	assert.Equal(t, transfer1.Amount, getTransfer.Amount)
}

func TestUpdateTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	arg := UpdateTransferParams{
		ID:     transfer1.ID,
		Amount: util.RandomMoney(),
	}
	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)
	assert.NoError(t, err)
	assert.Equal(t, transfer1.ID, transfer2.ID)

	assert.Equal(t, arg.Amount, transfer2.Amount)

}

func TestDeleteTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)
	assert.NoError(t, err)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	assert.Error(t, err)
	assert.Empty(t, transfer2)
}

func TestListTransfer(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}
	arg := ListTransferParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListTransfer(context.Background(), arg)
	assert.NoError(t, err)
	for _, account := range accounts {
		assert.NotEmpty(t, account)
	}
}

func TestListTransfersByFromAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		arg := CreateTransferParams{
			FromAccountID: int64(5),
			ToAccountID:   int64(6),
			Amount:        util.RandomMoney(),
		}
		_, err := testQueries.CreateTransfer(context.Background(), arg)
		assert.Nil(t, err)
	}
	arg := ListTransfersByFromAccountParams{
		FromAccountID: int64(5),
		Limit:         5,
		Offset:        5,
	}
	transfers, err := testQueries.ListTransfersByFromAccount(context.Background(), arg)
	assert.NoError(t, err)
	for _, transfer := range transfers {
		assert.NotEmpty(t, transfer)
		assert.Equal(t, int64(5), transfer.FromAccountID)
	}
}

func TestListTransfersByToAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		arg := CreateTransferParams{
			FromAccountID: int64(5),
			ToAccountID:   int64(6),
			Amount:        util.RandomMoney(),
		}
		_, err := testQueries.CreateTransfer(context.Background(), arg)
		assert.Nil(t, err)
	}
	arg := ListTransfersByToAccountParams{
		ToAccountID: int64(6),
		Limit:       5,
		Offset:      5,
	}
	transfers, err := testQueries.ListTransfersByToAccount(context.Background(), arg)
	assert.NoError(t, err)
	for _, transfer := range transfers {
		assert.NotEmpty(t, transfer)
		assert.Equal(t, int64(5), transfer.FromAccountID)
	}
}

func TestListTransfersByFromAndToAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		arg := CreateTransferParams{
			FromAccountID: int64(5),
			ToAccountID:   int64(6),
			Amount:        util.RandomMoney(),
		}
		_, err := testQueries.CreateTransfer(context.Background(), arg)
		assert.Nil(t, err)
	}
	arg := ListTransfersByFromAndToAccountParams{
		FromAccountID: int64(5),
		ToAccountID:   int64(6),
		Limit:         5,
		Offset:        5,
	}
	transfers, err := testQueries.ListTransfersByFromAndToAccount(context.Background(), arg)
	assert.NoError(t, err)
	for _, transfer := range transfers {
		assert.NotEmpty(t, transfer)
		assert.Equal(t, int64(5), transfer.FromAccountID)
		assert.Equal(t, int64(6), transfer.ToAccountID)
	}
}

func TestDeleteTransferByFromAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		arg := CreateTransferParams{
			FromAccountID: int64(5),
			ToAccountID:   int64(6),
			Amount:        util.RandomMoney(),
		}
		_, err := testQueries.CreateTransfer(context.Background(), arg)
		assert.Nil(t, err)
	}
	testQueries.DeleteTransfersByFromAccount(context.Background(), int64(5))
	arg := ListTransfersByFromAccountParams{
		FromAccountID: int64(5),
		Limit:         5,
		Offset:        5,
	}
	accounts, err := testQueries.ListTransfersByFromAccount(context.Background(), arg)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(accounts))
}

func TestDeleteTransferByToAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		arg := CreateTransferParams{
			FromAccountID: int64(5),
			ToAccountID:   int64(6),
			Amount:        util.RandomMoney(),
		}
		_, err := testQueries.CreateTransfer(context.Background(), arg)
		assert.Nil(t, err)
	}
	testQueries.DeleteTransfersByToAccount(context.Background(), int64(6))
	arg := ListTransfersByToAccountParams{
		ToAccountID: int64(6),
		Limit:       5,
		Offset:      5,
	}
	accounts, err := testQueries.ListTransfersByToAccount(context.Background(), arg)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(accounts))
}

func TestDeleteTransferByFromAndToAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		arg := CreateTransferParams{
			FromAccountID: int64(5),
			ToAccountID:   int64(6),
			Amount:        util.RandomMoney(),
		}
		_, err := testQueries.CreateTransfer(context.Background(), arg)
		assert.Nil(t, err)
	}
	arg := DeleteTransfersByFromAccountAndToAccountParams{
		FromAccountID: int64(5),
		ToAccountID:   int64(6),
	}
	testQueries.DeleteTransfersByFromAccountAndToAccount(context.Background(), arg)
	list := ListTransfersByFromAndToAccountParams{
		FromAccountID: int64(5),
		ToAccountID:   int64(6),
		Limit:         5,
		Offset:        5,
	}
	accounts, err := testQueries.ListTransfersByFromAndToAccount(context.Background(), list)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(accounts))
}
