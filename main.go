package main

import (
	"books-api/config/database"
	"books-api/server"
)

func main() {
	// Start DB
	database.StartDatabase()

	// Start server
	server := server.NewServer()

	server.Run()
}
