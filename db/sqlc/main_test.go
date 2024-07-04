package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/PyMarcus/go_sqlc/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot init load env")
		return
	}
	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
