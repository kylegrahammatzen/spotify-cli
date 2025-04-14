package repository

import (
	"database/sql"
	"spotify-cli/internal/models"
)

type ArtistRepository struct {
	db *sql.DB
}

func NewArtistRepository(db *sql.DB) *ArtistRepository {
	return &ArtistRepository{db: db}
}

func (r *ArtistRepository) GetByID(id int) (*models.Artist, error) {
	query := "SELECT id, name, genre FROM artists WHERE id = ?"
	row := r.db.QueryRow(query, id)

	artist := &models.Artist{}
	err := row.Scan(&artist.ID, &artist.Name, &artist.Genre)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (r *ArtistRepository) GetByName(name string) (*models.Artist, error) {
	query := "SELECT id, name, genre FROM artists WHERE name = ?"
	row := r.db.QueryRow(query, name)

	artist := &models.Artist{}
	err := row.Scan(&artist.ID, &artist.Name, &artist.Genre)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (r *ArtistRepository) List(limit, offset int) ([]*models.Artist, error) {
	query := "SELECT id, name, genre FROM artists LIMIT ? OFFSET ?"
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	artists := []*models.Artist{}
	for rows.Next() {
		artist := &models.Artist{}
		if err := rows.Scan(&artist.ID, &artist.Name, &artist.Genre); err != nil {
			return nil, err
		}
		artists = append(artists, artist)
	}

	return artists, nil
}

func (r *ArtistRepository) Create(artist *models.Artist) (int, error) {
	query := "INSERT INTO artists (name, genre) VALUES (?, ?)"
	result, err := r.db.Exec(query, artist.Name, artist.Genre)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return int(id), err
}

func (r *ArtistRepository) Update(artist *models.Artist) error {
	query := "UPDATE artists SET name = ?, genre = ? WHERE id = ?"
	_, err := r.db.Exec(query, artist.Name, artist.Genre, artist.ID)
	return err
}

func (r *ArtistRepository) Delete(id int) error {
	query := "DELETE FROM artists WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}
