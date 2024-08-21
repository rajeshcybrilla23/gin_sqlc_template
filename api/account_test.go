package api

import (
	"bytes"
	"encoding/json"
	mockdb "gin-template/db/mock"
	db "gin-template/db/sqlc"
	"gin-template/db/util"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestCreateAccountAPI(t *testing.T) {
	account := randomAccount("demo")

	arg := db.CreateAccountParams{
		Owner:    account.Owner,
		Currency: account.Currency,
		Balance:  0,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	// build stubbs
	store.EXPECT().
		CreateAccount(gomock.Any(), gomock.Eq(arg)).
		Times(1).
		Return(account, nil)

	// start server and send request
	server, _ := NewServer(store)
	recorder := httptest.NewRecorder()

	// Marshal body data to JSON
	body := gin.H{
		"owner":    account.Owner,
		"currency": account.Currency,
	}
	data, err := json.Marshal(body)
	require.NoError(t, err)

	url := "/accounts"
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	// check response
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchAccount(t, recorder.Body, account)
}

func randomAccount(owner string) db.Account {
	return db.Account{
		Owner:    owner,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotAccounts db.Account
	err = json.Unmarshal(data, &gotAccounts)
	require.NoError(t, err)
	require.Equal(t, account, gotAccounts)
}
