package db

import (
	"context"
	"os"
	"simplebank/util"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		panic(err)
	}
	testDB, err = pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		panic(err)
	}

	testQueries = New(testDB)

	code := m.Run()
	testDB.Close()
	os.Exit(code)
}
