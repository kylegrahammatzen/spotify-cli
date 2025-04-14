package db

import "spotify-cli/internal/models"

type Repository interface {
	Close() error

	// User operations
	GetUser(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) (int, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error

	// Artist operations
	GetArtist(id int) (*models.Artist, error)
	GetArtistByName(name string) (*models.Artist, error)
	ListArtists(limit, offset int) ([]*models.Artist, error)
	CreateArtist(artist *models.Artist) (int, error)
	UpdateArtist(artist *models.Artist) error
	DeleteArtist(id int) error

	// Album operations
	GetAlbum(id int) (*models.Album, error)
	GetAlbumsByArtist(artistID int) ([]*models.Album, error)
	ListAlbums(limit, offset int) ([]*models.Album, error)
	CreateAlbum(album *models.Album) (int, error)
	UpdateAlbum(album *models.Album) error
	DeleteAlbum(id int) error

	// Track operations
	GetTrack(id int) (*models.Track, error)
	GetTracksByAlbum(albumID int) ([]*models.Track, error)
	ListTracks(limit, offset int) ([]*models.Track, error)
	CreateTrack(track *models.Track) (int, error)
	UpdateTrack(track *models.Track) error
	DeleteTrack(id int) error

	// Playlist operations
	GetPlaylist(id int) (*models.Playlist, error)
	GetPlaylistsByUser(userID int) ([]*models.Playlist, error)
	ListPlaylists(limit, offset int) ([]*models.Playlist, error)
	CreatePlaylist(playlist *models.Playlist) (int, error)
	UpdatePlaylist(playlist *models.Playlist) error
	DeletePlaylist(id int) error

	// Item operations
	GetPlaylistTracks(playlistID int) ([]*models.Track, error)
	AddTrackToPlaylist(playlistID, trackID, position int) error
	RemoveTrackFromPlaylist(playlistID, trackID int) error

	// Setup operations
	SetupDB(sqlFilePath string) error
}
