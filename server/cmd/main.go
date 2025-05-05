package main

import (
	server "server-api/internal"
	"server-api/internal/database"
)

func main() {
	database.Start()
	server.Start()
}
