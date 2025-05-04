package main

import (
	"client-server-api/client"
	"client-server-api/server"
	"client-server-api/server/database"
)

func main() {
	database.Start()
	go server.Start()
	client.Start()
}
