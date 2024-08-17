package main

import (
	"context"
	"gin-template/api"
	db "gin-template/db/sqlc"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	const (
		dbDriver      = "postgres"
		dbSource      = "postgres://rajeshmanjunath:password@localhost:5432/gin_sqlc_template?sslmode=disable"
		serverAddress = "0.0.0.0:8080"
	)

	connPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(connPool)
	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot create server")
	}
}
