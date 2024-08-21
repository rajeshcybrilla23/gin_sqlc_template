package service

import (
	db "gin-template/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner"`
	Currency string `json:"currency" binding:"required,oneof=USD INR"`
}

func (server *DBStore) Create(ctx *gin.Context) (*db.Account, error) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}
	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
