package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testStore Store

const (
	dbDriver = "postgres"
	dbSource = "postgres://rajeshmanjunath:password@localhost:5432/gin_sqlc_template?sslmode=disable"
)

func TestMain(m *testing.M) {
	connPool, err := pgxpool.New(context.Background(), dbSource)
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

