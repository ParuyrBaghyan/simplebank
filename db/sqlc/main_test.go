package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:Paruyr20040324@localhost:5432/postgres?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can't connect to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
