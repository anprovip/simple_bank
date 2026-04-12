package main

import (
	"context"
	"simplebank/api"
	db "simplebank/db/sqlc"
	"simplebank/util"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		panic(err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		panic(err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		panic(err)
	}
}
