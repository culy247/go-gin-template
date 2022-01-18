package seeds

import (
	"log"

	"github.com/culy247/go-gin-template/config"
	"github.com/culy247/go-gin-template/models"
)

func Perfils() {
	db := config.Db()
	var Perfils []models.Perfil

	result := db.Find(&Perfils)
	if result.RowsAffected >= 4 {
		log.Printf("Seed (Perfils): Nothing to seed")
		//return
	}

	Perfils = []models.Perfil{
		{Description: "ADMIN"},
		{Description: "GUEST"},
	}
	for i := 0; i < 1000000; i++ {
		db.Create(&Perfils)
	}
	log.Printf("Seed (Perfils): Success")
}
