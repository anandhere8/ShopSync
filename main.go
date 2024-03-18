package main

import (
	"database/sql"
	"fmt"
	"log"

	server "github.com/anandhere8/ShopSync/cmd/shopsync"
	db "github.com/anandhere8/ShopSync/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@localhost:5432/shopsync?sslmode=disable"
	// dbSource = "user=root password=secret dbname=shopsync sslmode=disable"

	address = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	newQuery := db.New(conn)
	// fmt.Println(newQuery)
	server := server.NewServer(newQuery)
	fmt.Println(server)
	server.Start(address)
}
