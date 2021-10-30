package main

import (
	"github.com/zackwn/books-api/database"
	"github.com/zackwn/books-api/server"
)

func main() {
	database.Start()
	server := server.New()
	server.Run()
}
