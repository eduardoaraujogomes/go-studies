package main

import (
	"github.com/eduardoaraujogomes/go-studies/database"
	"github.com/eduardoaraujogomes/go-studies/routes"
)

func main() {
	database.DatabaseConnection()
	routes.HandleRequests()
}
