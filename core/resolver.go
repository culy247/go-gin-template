package core

import (
	"github.com/culy247/go-gin-template/config"
	"gorm.io/gorm"
)

type Resolver struct {
}

func (r *Resolver) Db() *gorm.DB {
	return config.Db()
}
