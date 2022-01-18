package database

import (
	"log"

	"github.com/culy247/go-gin-template/config"
	"github.com/culy247/go-gin-template/models"
)

func Migrate() {
	log.Printf("Migrate: Start")
	db := config.Db()
	db.AutoMigrate(
		&models.User{},
	)
	log.Printf("Migrate: Success")
}
