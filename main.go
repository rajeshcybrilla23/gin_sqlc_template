package main

import (
	"context"
	"gin-template/api"
	db "gin-template/db/sqlc"
	"gin-template/db/util"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".") // . mean current folder
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(connPool)
	api.NewServer(store)
	router := api.SetupRouter()

	err = start(config.HTTPServerAddress, router)
	if err != nil {
		log.Fatal("cannot create server")
	}
}

// Start runs the HTTP server on a specific address.
func start(address string, router *gin.Engine) error {
	return router.Run(address)
}
