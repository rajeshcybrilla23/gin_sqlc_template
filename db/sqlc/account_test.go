package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx"
	"github.com/stretchr/testify/require"
)

func setUp() {
	// SETUP METHOD WHICH IS REQUIRED TO RUN FOR EACH TEST METHOD
	// your code here
	//https://stackoverflow.com/questions/23729790/how-can-i-do-test-setup-using-the-testing-package-in-go/65428147#65428147
}

func tearDown(ctx context.Context) {
	// TEAR DOWN METHOD WHICH IS REQUIRED TO RUN FOR EACH TEST METHOD
	// your code here
	testStore.DeleteAll(context.Background())
}

var RunTest = CreateForEach(setUp, tearDown)

func createRandomAccount(t *testing.T) Account {

	arg := CreateAccountParams{
		Owner:    "tom",
		Balance:  100,
		Currency: "USD",
	}

	account, err := testStore.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	RunTest(func() {
		createRandomAccount(t)
	})
}

func TestGetAccount(t *testing.T) {
	RunTest(func() {
		account1 := createRandomAccount(t)
		account2, err := testStore.GetAccount(context.Background(), account1.ID)
		require.NoError(t, err)
		require.NotEmpty(t, account2)

		require.Equal(t, account1.ID, account2.ID)
		require.Equal(t, account1.Owner, account2.Owner)
		require.Equal(t, account1.Balance, account2.Balance)
		require.Equal(t, account1.Currency, account2.Currency)
		require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
	})
}

func TestUpdateAccount(t *testing.T) {
	RunTest(func() {
		account1 := createRandomAccount(t)

		arg := UpdateAccountParams{
			ID:      account1.ID,
			Balance: 200,
		}

		account2, err := testStore.UpdateAccount(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, account2)

		require.Equal(t, account1.ID, account2.ID)
		require.Equal(t, account1.Owner, account2.Owner)
		require.Equal(t, arg.Balance, account2.Balance)
		require.Equal(t, account1.Currency, account2.Currency)
		require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
	})
}

func TestDeleteAccount(t *testing.T) {
	RunTest(func() {
		account1 := createRandomAccount(t)
		err := testStore.DeleteAccount(context.Background(), account1.ID)
		require.NoError(t, err)

		account2, err := testStore.GetAccount(context.Background(), account1.ID)
		require.Error(t, err)
		require.EqualError(t, err, pgx.ErrNoRows.Error())
		require.Empty(t, account2)
	})
}

func TestListAccounts(t *testing.T) {
	RunTest(func() {
		var lastAccount Account
		for i := 0; i < 10; i++ {
			lastAccount = createRandomAccount(t)
		}

		arg := ListAccountsParams{
			Owner:  lastAccount.Owner,
			Limit:  5,
			Offset: 0,
		}

		accounts, err := testStore.ListAccounts(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, accounts)

		for _, account := range accounts {
			require.NotEmpty(t, account)
			require.Equal(t, lastAccount.Owner, account.Owner)
		}
	})
}

func TestTxCreateAccount(t *testing.T) {
	RunTest(func() {
		arg := CreateAccountParams{
			Owner:    "tom",
			Balance:  100,
			Currency: "USD",
		}

		params := CreateAccountTxParams{
			CreateAccountParams: arg,
		}
		result, err := testStore.CreateAccountTx(context.Background(), params)

		require.NoError(t, err)
		require.NotEmpty(t, result)
	})
}
