package repository

import (
	"context"
	db "gin-template/db/sqlc"
)

type AccountRpository interface {
	CreateAccount(ctx context.Context, arg db.CreateAccountParams) (db.Account, error)
}

type AccountRpositoryImp struct {
	store db.Store
}

func AccountRepositoryImpInit(store db.Store) *AccountRpositoryImp {
	return &AccountRpositoryImp{
		store: store,
	}
}

func (accImpl AccountRpositoryImp) CreateAccount(ctx context.Context, arg db.CreateAccountParams) (db.Account, error) {

	account, err := accImpl.store.CreateAccount(ctx, arg)
	if err != nil {

		return db.Account{}, err
	}

	return account, nil
}
