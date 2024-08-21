package api

import (
	db "gin-template/db/sqlc"
)

type Server struct {
	store db.Store
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(store db.Store) *Server {

	server := &Server{
		store: store,
	}

	return server
}
