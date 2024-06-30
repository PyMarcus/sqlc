package main

import (
	"database/sql"
	_ "github.com/lib/pq"

	"log"

	"github.com/PyMarcus/go_sqlc/api"
	db "github.com/PyMarcus/go_sqlc/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbStr    = "postgresql://myuser:secret@localhost:5432/db?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbStr)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	server := api.NewServer(db.NewStore(conn))
	addr := ":8080"
	log.Println("Running...", addr)
	server.Start(addr)
}
