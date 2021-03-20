package main

import (
	"database/sql"
	"log"
)

func main() {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	store := NewStore(conn)
	router := NewRouter(store)
	server := NewServer(config, store, router)

	server.Run()
}
