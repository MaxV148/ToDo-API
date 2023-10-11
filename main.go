package main

import (
	"CheckToDoAPI/api"
	db "CheckToDoAPI/db/sqlc"
	"CheckToDoAPI/utils"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalln("cannot load config: ", err)
	}

	log.Println("CURRENT DATABASE: ", config.DBSourceProd)
	conn, err := sql.Open(config.DBDriver, config.DBSourceProd)
	if err != nil {
		log.Fatalln("cannot connect to DBMS: ", err)
	}

	queries := db.New(conn)
	server := api.NewServer(queries, config)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalln("cannot start Server: ", err)

	}
}
