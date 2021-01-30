package main

import (
	"gubler/db"
	"gubler/server"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	var err error
	godotenv.Load()

	err = db.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	err = server.RunServer()

	if err != nil {
		log.Fatal(err)
	}

	db.DisconnectDB()

}
