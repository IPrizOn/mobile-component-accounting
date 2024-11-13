package main

import (
	"mobile/client"
	"mobile/database"
)

func main() {
	client.Start()

	database.InitDatabase()
}
