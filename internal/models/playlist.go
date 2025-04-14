package models

import "time"

type Playlist struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`

	User   *User    `json:"user,omitempty"`
	Tracks []*Track `json:"tracks,omitempty"`
}
