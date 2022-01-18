package database

import (
	"log"

	"github.com/culy247/go-gin-template/database/seeds"
)

func Seeder() {
	log.Printf("Seed: Start")
	seeds.Perfils()
	log.Printf("Seed: Success")
}
