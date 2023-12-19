package handler

import (
	"github.com/nbrkwnm/golang-webapi/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeHandler() {
	db = config.GetPostgreSQL()
}
