package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Chengxufeng1994/go-ddd/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPgGorm(cfg *config.Postgres) (*gorm.DB, error) {
	_default := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		})

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName)
	pgCfg := postgres.Config{
		DSN:                  dsn, // DSN data source name
		PreferSimpleProtocol: false,
	}
	db, err := gorm.Open(postgres.New(pgCfg),
		&gorm.Config{
			Logger: _default,
		})
	if err != nil {
		return nil, err
	}

	return db, nil
}
