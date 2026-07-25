package sqlite

import (
	"7DL/config"

	sqlitedriver "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteDB struct {
	Config config.SQLiteConfig
	DB     *gorm.DB
}

func New(config config.SQLiteConfig) (*SQLiteDB, error) {
	db, err := gorm.Open(sqlitedriver.Open(config.Path), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &SQLiteDB{
		Config: config,
		DB:     db,
	}, nil
}
