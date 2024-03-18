package db

import (
	"database/sql"
	"log"
	"sync"

	db "github.com/anandhere8/ShopSync/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:5432/shopsync?sslmode=disable"
	// dbSource = "user=root password=secret dbname=shopsync sslmode=disable"

	address = "0.0.0.0:8080"
)

var (
	err        error
	once       sync.Once
	dbInstance *db.Queries
)

func GetDBInstance() (*db.Queries, error) {

	once.Do(func() {
		conn, err := sql.Open(dbDriver, dbSource)
		if err != nil {
			log.Fatal("Cannot connect to database", err)
			return
		}
		dbInstance = db.New(conn)
	})

	return dbInstance, err
}
