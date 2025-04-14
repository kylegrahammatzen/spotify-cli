package models

import "time"

type Item struct {
	PlaylistID int       `json:"playlist_id"`
	TrackID    int       `json:"track_id"`
	Position   int       `json:"position"`
	AddedAt    time.Time `json:"added_at"`

	Track *Track `json:"track,omitempty"`
}
