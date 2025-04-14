package repository

import (
	"database/sql"
	"spotify-cli/internal/models"
)

type AlbumRepository struct {
	db *sql.DB
}

func NewAlbumRepository(db *sql.DB) *AlbumRepository {
	return &AlbumRepository{db: db}
}

func (r *AlbumRepository) GetByID(id int) (*models.Album, error) {
	query := "SELECT id, title, artist_id, release_date FROM albums WHERE id = ?"
	row := r.db.QueryRow(query, id)

	album := &models.Album{}
	err := row.Scan(&album.ID, &album.Title, &album.ArtistID, &album.ReleaseDate)
	if err != nil {
		return nil, err
	}

	return album, nil
}

func (r *AlbumRepository) GetByArtist(artistID int) ([]*models.Album, error) {
	query := "SELECT id, title, artist_id, release_date FROM albums WHERE artist_id = ?"
	rows, err := r.db.Query(query, artistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	albums := []*models.Album{}
	for rows.Next() {
		album := &models.Album{}
		if err := rows.Scan(&album.ID, &album.Title, &album.ArtistID, &album.ReleaseDate); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}

	return albums, nil
}

func (r *AlbumRepository) List(limit, offset int) ([]*models.Album, error) {
	query := "SELECT id, title, artist_id, release_date FROM albums LIMIT ? OFFSET ?"
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	albums := []*models.Album{}
	for rows.Next() {
		album := &models.Album{}
		if err := rows.Scan(&album.ID, &album.Title, &album.ArtistID, &album.ReleaseDate); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}

	return albums, nil
}

func (r *AlbumRepository) Create(album *models.Album) (int, error) {
	query := "INSERT INTO albums (title, artist_id, release_date) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, album.Title, album.ArtistID, album.ReleaseDate)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

func (r *AlbumRepository) Update(album *models.Album) error {
	query := "UPDATE albums SET title = ?, artist_id = ?, release_date = ? WHERE id = ?"
	_, err := r.db.Exec(query, album.Title, album.ArtistID, album.ReleaseDate, album.ID)
	return err
}

func (r *AlbumRepository) Delete(id int) error {
	query := "DELETE FROM albums WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}
