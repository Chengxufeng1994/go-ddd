package database

import (
	"fmt"

	"github.com/Chengxufeng1994/go-ddd/config"
	"gorm.io/gorm"
)

func New(cfg *config.Persistence) (*gorm.DB, error) {
	switch cfg.Type {
	case "postgres":
		return NewPgGorm(&cfg.Postgres)
	default:
		return nil, fmt.Errorf("unknown database type")
	}
}
