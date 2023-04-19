package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ttanik/simplebank/util"
)

func createRandomEntry(t *testing.T) Entry {
	arg := CreateEntryParams{
		AccountID: util.RandomInt(1, 2),
		Amount:    util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	assert.Nil(t, err)
	assert.Equal(t, arg.AccountID, entry.AccountID)
	assert.Equal(t, arg.Amount, entry.Amount)
	return entry
}
func TestCreateEntry(t *testing.T) {
	_ = createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	account1 := createRandomEntry(t)
	getAccount, err := testQueries.GetEntry(context.Background(), account1.ID)
	assert.NoError(t, err)
	assert.Equal(t, account1.ID, getAccount.ID)
	assert.Equal(t, account1.Amount, getAccount.Amount)
}

func TestUpdateEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	arg := UpdateEntryParams{
		ID:     entry1.ID,
		Amount: util.RandomMoney(),
	}
	entry2, err := testQueries.UpdateEntry(context.Background(), arg)
	assert.NoError(t, err)
	assert.Equal(t, entry1.ID, entry2.ID)

	assert.Equal(t, arg.Amount, entry2.Amount)

}

func TestDeleteEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	assert.NoError(t, err)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	assert.Error(t, err)
	assert.Empty(t, entry2)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}
	arg := ListEntryParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListEntry(context.Background(), arg)
	assert.NoError(t, err)
	for _, account := range accounts {
		assert.NotEmpty(t, account)
	}
}

func TestListEntriesByAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		arg := CreateEntryParams{
			AccountID: int64(2),
			Amount:    util.RandomMoney(),
		}
		_, err := testQueries.CreateEntry(context.Background(), arg)
		assert.Nil(t, err)
	}
	arg := ListEntriesByAccountParams{
		AccountID: int64(2),
		Limit:     5,
		Offset:    5,
	}
	accounts, err := testQueries.ListEntriesByAccount(context.Background(), arg)
	assert.NoError(t, err)
	for _, account := range accounts {
		assert.NotEmpty(t, account)
		assert.Equal(t, int64(2), account.AccountID)
	}
}

func TestDeleteEntryByAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		arg := CreateEntryParams{
			AccountID: int64(2),
			Amount:    util.RandomMoney(),
		}
		_, err := testQueries.CreateEntry(context.Background(), arg)
		assert.Nil(t, err)
	}
	testQueries.DeleteEntriesByAccount(context.Background(), int64(2))
	arg := ListEntriesByAccountParams{
		AccountID: int64(2),
		Limit:     5,
		Offset:    5,
	}
	accounts, err := testQueries.ListEntriesByAccount(context.Background(), arg)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(accounts))
}
