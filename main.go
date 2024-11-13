package main

import (
	"mobile/client"
	"mobile/database"
)

func main() {
	database.InitDatabase()

	client.Start()
}
