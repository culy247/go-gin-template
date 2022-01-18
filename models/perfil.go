package models

import (
	"github.com/culy247/go-gin-template/core"
)

type Perfil struct {
	Name        string `json:"name" gorm:"index;not null"`
	Description string `json:"description" gorm:"not null"`
	core.Model
}
