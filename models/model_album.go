package models

import (
	"time"

	"gorm.io/gorm"
)

type Album struct {
	gorm.Model
	Title  string
	Artist string
	Price  float64
}

type AlbumResponse struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Title     string    `json:"title"`
	Artist    string    `json:"artist"`
	Price     float64   `json:"price"`
}

type AlbumRequest struct {
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
