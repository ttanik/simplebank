package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTXResult)

	for i := 0; i < n; i++ {
		go func() {

			result, err := store.TransferTx(context.Background(), TransferTXParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}
	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		assert.NoError(t, err)

		result := <-results
		assert.NotEmpty(t, result)

		transfer := result.Transfer
		assert.NotEmpty(t, transfer)
		assert.Equal(t, account1.ID, transfer.FromAccountID)
		assert.Equal(t, account1.ID, transfer.FromAccountID)
		assert.Equal(t, account2.ID, transfer.ToAccountID)
		assert.Equal(t, amount, transfer.Amount)
		assert.NotZero(t, transfer.ID)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		assert.NoError(t, err)

		fromEntry := result.FromEntry
		assert.NotEmpty(t, fromEntry)
		assert.Equal(t, account1.ID, fromEntry.AccountID)
		assert.Equal(t, -amount, fromEntry.Amount)
		assert.NotZero(t, fromEntry.ID)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		assert.NoError(t, err)

		toEntry := result.ToEntry
		assert.NotEmpty(t, toEntry)
		assert.Equal(t, account2.ID, toEntry.AccountID)
		assert.Equal(t, amount, toEntry.Amount)
		assert.NotZero(t, toEntry.ID)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		assert.NoError(t, err)

		fromAccount := result.FromAccount
		assert.NotEmpty(t, fromAccount)
		assert.Equal(t, account1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		assert.NotEmpty(t, toAccount)
		assert.Equal(t, account2.ID, toAccount.ID)

		//check accounts balance
		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		assert.Equal(t, diff1, diff2)
		assert.True(t, diff1 > 0)
		assert.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		assert.True(t, k >= 1 && k <= n)
		assert.NotContains(t, existed, k)
		existed[k] = true
	}

	//check final balances
	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	assert.NoError(t, err)
	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	assert.NoError(t, err)

	assert.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
	assert.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)
}

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 10
	amount := int64(10)

	errs := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}
		go func() {
			_, err := store.TransferTx(context.Background(), TransferTXParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})
			errs <- err

		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		assert.NoError(t, err)
	}

	//check final balances
	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	assert.NoError(t, err)
	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	assert.NoError(t, err)

	assert.Equal(t, account1.Balance, updatedAccount1.Balance)
	assert.Equal(t, account2.Balance, updatedAccount2.Balance)
}
