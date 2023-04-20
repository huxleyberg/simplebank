package main

import (
	"database/sql"
	"github.com/huxleyberg/simplebank/api"
	db "github.com/huxleyberg/simplebank/db/sqlc"
	"github.com/huxleyberg/simplebank/util"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalln("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatalln("cannot start server", err)
	}
}
