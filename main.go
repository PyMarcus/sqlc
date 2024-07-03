package main

import (
	"database/sql"

	_ "github.com/lib/pq"

	"log"

	"github.com/PyMarcus/go_sqlc/api"
	db "github.com/PyMarcus/go_sqlc/db/sqlc"
	"github.com/PyMarcus/go_sqlc/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Println(err)
		return
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	server := api.NewServer(db.NewStore(conn))
	addr := config.ServerAddr
	log.Println("Running...", addr)
	server.Start(addr)
}
