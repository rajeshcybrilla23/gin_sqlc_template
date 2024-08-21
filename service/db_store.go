package service

import (
	db "gin-template/db/sqlc"
)

type DBStore struct {
	store db.Store
}
