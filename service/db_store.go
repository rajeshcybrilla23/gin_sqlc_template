package service

import (
	db "gin-template/db/sqlc"
)

type DBStore struct {
	store db.Store
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(store db.Store) *DBStore {

	server := &DBStore{
		store: store,
	}

	return server
}
