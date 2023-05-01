package main

import (
	"database/sql"
	"log"

	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/lib/pq"
	"github.com/ttanik/simplebank/api"
	db "github.com/ttanik/simplebank/db/sqlc"
	"github.com/ttanik/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configurations: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server")
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
