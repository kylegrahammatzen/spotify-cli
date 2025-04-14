package models

import "time"

type Album struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	ArtistID    int       `json:"artist_id"`
	ReleaseDate time.Time `json:"release_date,omitempty"`

	Artist *Artist `json:"artist,omitempty"`
}
