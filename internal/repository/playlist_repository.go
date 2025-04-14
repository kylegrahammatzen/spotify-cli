package repository

import (
	"database/sql"
	"spotify-cli/internal/models"
)

type PlaylistRepository struct {
	db *sql.DB
}

func NewPlaylistRepository(db *sql.DB) *PlaylistRepository {
	return &PlaylistRepository{db: db}
}

func (r *PlaylistRepository) GetByID(id int) (*models.Playlist, error) {
	query := "SELECT id, title, user_id, created_at FROM playlists WHERE id = ?"
	row := r.db.QueryRow(query, id)

	playlist := &models.Playlist{}
	err := row.Scan(&playlist.ID, &playlist.Title, &playlist.UserID, &playlist.CreatedAt)
	if err != nil {
		return nil, err
	}

	return playlist, nil
}

func (r *PlaylistRepository) GetByUser(userID int) ([]*models.Playlist, error) {
	query := "SELECT id, title, user_id, created_at FROM playlists WHERE user_id = ?"
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	playlists := []*models.Playlist{}
	for rows.Next() {
		playlist := &models.Playlist{}
		if err := rows.Scan(&playlist.ID, &playlist.Title, &playlist.UserID, &playlist.CreatedAt); err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

func (r *PlaylistRepository) List(limit, offset int) ([]*models.Playlist, error) {
	query := "SELECT id, title, user_id, created_at FROM playlists LIMIT ? OFFSET ?"
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	playlists := []*models.Playlist{}
	for rows.Next() {
		playlist := &models.Playlist{}
		if err := rows.Scan(&playlist.ID, &playlist.Title, &playlist.UserID, &playlist.CreatedAt); err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}

	return playlists, nil
}

func (r *PlaylistRepository) Create(playlist *models.Playlist) (int, error) {
	query := "INSERT INTO playlists (title, user_id) VALUES (?, ?)"
	result, err := r.db.Exec(query, playlist.Title, playlist.UserID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

func (r *PlaylistRepository) Update(playlist *models.Playlist) error {
	query := "UPDATE playlists SET title = ? WHERE id = ?"
	_, err := r.db.Exec(query, playlist.Title, playlist.ID)
	return err
}

func (r *PlaylistRepository) Delete(id int) error {
	query := "DELETE FROM playlists WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PlaylistRepository) GetTracks(playlistID int) ([]*models.Track, error) {
	query := `
		SELECT t.id, t.title, t.album_id, t.duration 
		FROM tracks t
		JOIN items i ON t.id = i.track_id
		WHERE i.playlist_id = ?
		ORDER BY i.position
	`
	rows, err := r.db.Query(query, playlistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tracks := []*models.Track{}
	for rows.Next() {
		track := &models.Track{}
		if err := rows.Scan(&track.ID, &track.Title, &track.AlbumID, &track.Duration); err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (r *PlaylistRepository) AddTrack(playlistID, trackID, position int) error {
	query := "INSERT INTO items (playlist_id, track_id, position) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, playlistID, trackID, position)
	return err
}

func (r *PlaylistRepository) RemoveTrack(playlistID, trackID int) error {
	query := "DELETE FROM items WHERE playlist_id = ? AND track_id = ?"
	_, err := r.db.Exec(query, playlistID, trackID)
	return err
}
