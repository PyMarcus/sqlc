package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDb *sql.DB

const (
	dbDriver = "postgres"
	dbStr    = "postgresql://myuser:secret@localhost:5432/db?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error 
	
	testDb, err = sql.Open(dbDriver, dbStr)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
