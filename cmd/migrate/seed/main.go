package main

import (
	"first_mod/internal/db"
	"first_mod/internal/env"
	"first_mod/internal/store"
	"log"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable")
	conn, err := db.New(addr, 30, 30, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	store := store.NewStorage(conn)
	db.Seed(store)
}
