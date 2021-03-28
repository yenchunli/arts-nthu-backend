package main

import (
	"database/sql"
	"fmt"
	db "github.com/yenchunli/go-nthu-artscenter-server/db"
	"log"
)

func main() {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	fmt.Println(psqlInfo)
	conn, err := sql.Open(config.DBDriver, psqlInfo)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	store := db.NewDB(conn)
	server, _ := NewServer(config, store)

	server.Run()
}
