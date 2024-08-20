package db

import (
	"context"
	"gin-template/db/util"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("./../..") // . mean current folder
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testStore = NewStore(connPool)

	// inform os that connnection was successfulr nad terminate the test case
	os.Exit(m.Run())
}

// setup and tear down
func CreateForEach(setUp func(), tearDown func(ctx context.Context)) func(func()) {
	ctx := context.Background()
	return func(testFunc func()) {
		setUp()
		testFunc()
		tearDown(ctx)
	}
}
