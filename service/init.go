package service

import (
	db "gin-template/db/sqlc"
)

type Services struct {
	AccountService AccountService
}

func InitServices(store db.Store) Services {
	return Services{
		AccountService: AccountService{
			DBStore: DBStore{
				store: store,
			},
		},
	}
}
