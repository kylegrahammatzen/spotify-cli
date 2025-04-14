package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"spotify-cli/internal/models"
)

func NewPlaylistCommand(repos *Repositories) *cli.Command {
	return &cli.Command{
		Name:  "playlists",
		Usage: "Manage playlists",
		Subcommands: []*cli.Command{
			ListPlaylistsCommand(repos),
			CreatePlaylistCommand(repos),
			PlaylistTracksCommand(repos),
			AddTrackToPlaylistCommand(repos),
			RemoveTrackFromPlaylistCommand(repos),
		},
	}
}

func ListPlaylistsCommand(repos *Repositories) *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "List playlists",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "limit",
				Value: 10,
				Usage: "Number of results to show",
			},
			&cli.IntFlag{
				Name:  "offset",
				Value: 0,
				Usage: "Result offset for pagination",
			},
			&cli.IntFlag{
				Name:  "user",
				Usage: "Filter by user ID",
			},
		},
		Action: func(c *cli.Context) error {
			limit := c.Int("limit")
			offset := c.Int("offset")
			userID := c.Int("user")

			var playlists []*models.Playlist
			var err error

			if userID > 0 {
				playlists, err = repos.PlaylistRepo.GetByUser(userID)
			} else {
				playlists, err = repos.PlaylistRepo.List(limit, offset)
			}

			if err != nil {
				return err
			}

			fmt.Printf("Found %d playlists:\n", len(playlists))
			for _, playlist := range playlists {
				fmt.Printf("ID: %d, Title: %s, User ID: %d, Created: %s\n",
					playlist.ID, playlist.Title, playlist.UserID, playlist.CreatedAt.Format("2006-01-02"))
			}

			return nil
		},
	}
}

func CreatePlaylistCommand(repos *Repositories) *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Create a new playlist",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "title",
				Aliases:  []string{"t"},
				Usage:    "Playlist title",
				Required: true,
			},
			&cli.IntFlag{
				Name:     "user",
				Aliases:  []string{"u"},
				Usage:    "User ID",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			title := c.String("title")
			userID := c.Int("user")

			playlist := &models.Playlist{
				Title:  title,
				UserID: userID,
			}

			id, err := repos.PlaylistRepo.Create(playlist)
			if err != nil {
				return err
			}

			fmt.Printf("Playlist created with ID: %d\n", id)
			return nil
		},
	}
}

func PlaylistTracksCommand(repos *Repositories) *cli.Command {
	return &cli.Command{
		Name:  "tracks",
		Usage: "List tracks in a playlist",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "id",
				Usage:    "Playlist ID",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			playlistID := c.Int("id")

			tracks, err := repos.PlaylistRepo.GetTracks(playlistID)
			if err != nil {
				return err
			}

			playlist, err := repos.PlaylistRepo.GetByID(playlistID)
			if err != nil {
				return err
			}

			fmt.Printf("Tracks in playlist '%s':\n", playlist.Title)
			for i, track := range tracks {
				duration := formatDuration(track.Duration)
				fmt.Printf("%d. %s (%s)\n", i+1, track.Title, duration)
			}

			return nil
		},
	}
}

func AddTrackToPlaylistCommand(repos *Repositories) *cli.Command {
	return &cli.Command{
		Name:  "add-track",
		Usage: "Add a track to a playlist",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "playlist",
				Aliases:  []string{"p"},
				Usage:    "Playlist ID",
				Required: true,
			},
			&cli.IntFlag{
				Name:     "track",
				Aliases:  []string{"t"},
				Usage:    "Track ID",
				Required: true,
			},
			&cli.IntFlag{
				Name:  "position",
				Usage: "Position in playlist (defaults to end)",
			},
		},
		Action: func(c *cli.Context) error {
			playlistID := c.Int("playlist")
			trackID := c.Int("track")
			position := c.Int("position")

			if position <= 0 {
				tracks, err := repos.PlaylistRepo.GetTracks(playlistID)
				if err != nil {
					return err
				}
				position = len(tracks) + 1
			}

			err := repos.PlaylistRepo.AddTrack(playlistID, trackID, position)
			if err != nil {
				return err
			}

			fmt.Printf("Track added to playlist at position %d\n", position)
			return nil
		},
	}
}

func RemoveTrackFromPlaylistCommand(repos *Repositories) *cli.Command {
	return &cli.Command{
		Name:  "remove-track",
		Usage: "Remove a track from a playlist",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "playlist",
				Aliases:  []string{"p"},
				Usage:    "Playlist ID",
				Required: true,
			},
			&cli.IntFlag{
				Name:     "track",
				Aliases:  []string{"t"},
				Usage:    "Track ID",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			playlistID := c.Int("playlist")
			trackID := c.Int("track")

			err := repos.PlaylistRepo.RemoveTrack(playlistID, trackID)
			if err != nil {
				return err
			}

			fmt.Println("Track removed from playlist")
			return nil
		},
	}
}
