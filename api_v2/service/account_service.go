package service

import (
	"gin-template/api_v2/repository"

	db "gin-template/db/sqlc"

	"github.com/gin-gonic/gin"
)

type AccountService interface {
	createAccount(ctx *gin.Context)
}

type AccountServiceImpl struct {
	accRepo repository.AccountRpositoryImp
}

type createAccountRequest struct {
	Owner    string `json:"owner"`
	Currency string `json:"currency" binding:"required,oneof=USD INR"`
}

func AccountServiceImplInit(accRepo repository.AccountRpositoryImp) AccountServiceImpl {
	return AccountServiceImpl{
		accRepo: accRepo,
	}
}

func (accServ AccountServiceImpl) createAccount(ctx *gin.Context) (db.Account, error) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return db.Account{}, err
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := accServ.accRepo.CreateAccount(ctx, arg)
	if err != nil {
		return db.Account{}, err
	}

	return account, nil
}
