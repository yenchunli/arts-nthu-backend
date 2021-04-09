package main

import (
	"database/sql"
	"fmt"
	db "github.com/yenchunli/arts-nthu-backend/db"
	"github.com/yenchunli/arts-nthu-backend/util"
	"github.com/yenchunli/arts-nthu-backend/server"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	
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
	server, _ := server.NewServer(config, store)

	server.Run()
}
