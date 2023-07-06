package main

import (
	"autoGarage/api"
	db "autoGarage/db/sqlc"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@postgres:5432/garage_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Println("cannot connect to db")
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	if err := server.Run(serverAddress); err != nil {
		log.Printf("server cannot run at address: %v", err)
	}
}
