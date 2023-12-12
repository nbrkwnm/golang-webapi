package config

import "gorm.io/gorm"

var (
	db  *gorm.DB
	err error
)

func Init() error {

	db, err = InitializePostgreSQL()

	if err != nil {
		return err
	}

	return nil
}

func GetPostgreSQL() *gorm.DB {
	return db
}
