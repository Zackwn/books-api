package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/zackwn/books-api/database"
	"github.com/zackwn/books-api/server"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	database.Start()
	server := server.New()
	server.Run()
}
