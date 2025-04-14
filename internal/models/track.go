package models

type Track struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	AlbumID  int    `json:"album_id"`
	Duration int    `json:"duration"`

	Album *Album `json:"album,omitempty"`
}
