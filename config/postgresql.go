package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dns = "host=localhost port=5432 user=postgres password=postgres123 dbname=gorm sslmode=disable"
)

func InitializePostgreSQL() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
