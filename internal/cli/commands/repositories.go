package commands

import (
	"spotify-cli/internal/repository"
)

type Repositories struct {
	UserRepo     *repository.UserRepository
	ArtistRepo   *repository.ArtistRepository
	AlbumRepo    *repository.AlbumRepository
	TrackRepo    *repository.TrackRepository
	PlaylistRepo *repository.PlaylistRepository
}
