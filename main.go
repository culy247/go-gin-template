package main

import (
	"github.com/culy247/go-gin-template/database"
	"github.com/culy247/go-gin-template/routes"
)

func main() {
	database.Migrate()
	database.Seeder()
	routes.Run()
}
