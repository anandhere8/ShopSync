package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:5432/shopsync?sslmode=disable"
	// dbSource = "user=root password=secret dbname=shopsync sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Failed to connect to database", err)

	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
