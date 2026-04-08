package db

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbDriver        = "pgx"
	defaultDBSource = "postgresql://root:213077@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	testDB, err = pgxpool.New(context.Background(), defaultDBSource)
	if err != nil {
		panic(err)
	}

	testQueries = New(testDB)

	code := m.Run()
	testDB.Close()
	os.Exit(code)
}
