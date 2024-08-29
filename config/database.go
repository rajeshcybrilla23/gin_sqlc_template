package config

import (
	"context"
	db "gin-template/db/sqlc"
	"gin-template/db/util"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToDB(env util.Config) db.Store {
	connPool, err := pgxpool.New(context.Background(), env.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	return db.NewStore(connPool)
}
