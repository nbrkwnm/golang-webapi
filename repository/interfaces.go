package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/nbrkwnm/golang-webapi/models"
)

type AlbumPersistence interface {
	GetAlbumById(context.Context, uuid.UUID) (*models.Album, error)
	GetAlbums(context.Context) ([]*models.Album, error)
	CreateAlbum(context.Context, *models.Album) error
	UpdateAlbum(context.Context, *models.Album) error
	DeleteAlbum(context.Context, uuid.UUID) error
}
