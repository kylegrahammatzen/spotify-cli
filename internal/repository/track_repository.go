package repository

import (
	"database/sql"
	"spotify-cli/internal/models"
)

type TrackRepository struct {
	db *sql.DB
}

func NewTrackRepository(db *sql.DB) *TrackRepository {
	return &TrackRepository{db: db}
}

func (r *TrackRepository) GetByID(id int) (*models.Track, error) {
	query := "SELECT id, title, album_id, duration FROM tracks WHERE id = ?"
	row := r.db.QueryRow(query, id)

	track := &models.Track{}
	err := row.Scan(&track.ID, &track.Title, &track.AlbumID, &track.Duration)
	if err != nil {
		return nil, err
	}

	return track, nil
}

func (r *TrackRepository) GetByAlbum(albumID int) ([]*models.Track, error) {
	query := "SELECT id, title, album_id, duration FROM tracks WHERE album_id = ?"
	rows, err := r.db.Query(query, albumID)
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

func (r *TrackRepository) List(limit, offset int) ([]*models.Track, error) {
	query := "SELECT id, title, album_id, duration FROM tracks LIMIT ? OFFSET ?"
	rows, err := r.db.Query(query, limit, offset)
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

func (r *TrackRepository) Create(track *models.Track) (int, error) {
	query := "INSERT INTO tracks (title, album_id, duration) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, track.Title, track.AlbumID, track.Duration)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

func (r *TrackRepository) Update(track *models.Track) error {
	query := "UPDATE tracks SET title = ?, album_id = ?, duration = ? WHERE id = ?"
	_, err := r.db.Exec(query, track.Title, track.AlbumID, track.Duration, track.ID)
	return err
}

func (r *TrackRepository) Delete(id int) error {
	query := "DELETE FROM tracks WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}
