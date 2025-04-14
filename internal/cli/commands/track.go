package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"spotify-cli/internal/models"
)

func NewTrackCommand(repos *Repositories) *cli.Command {
	return &cli.Command{
		Name:  "tracks",
		Usage: "Manage tracks",
		Subcommands: []*cli.Command{
			ListTracksCommand(repos),
			AddTrackCommand(repos),
		},
	}
}

func ListTracksCommand(repos *Repositories) *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "List tracks",
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
				Name:  "album",
				Usage: "Filter by album ID",
			},
		},
		Action: func(c *cli.Context) error {
			limit := c.Int("limit")
			offset := c.Int("offset")
			albumID := c.Int("album")

			var tracks []*models.Track
			var err error

			if albumID > 0 {
				tracks, err = repos.TrackRepo.GetByAlbum(albumID)
			} else {
				tracks, err = repos.TrackRepo.List(limit, offset)
			}

			if err != nil {
				return err
			}

			fmt.Printf("Found %d tracks:\n", len(tracks))
			for _, track := range tracks {
				duration := formatDuration(track.Duration)
				fmt.Printf("ID: %d, Title: %s, Album ID: %d, Duration: %s\n",
					track.ID, track.Title, track.AlbumID, duration)
			}

			return nil
		},
	}
}

func AddTrackCommand(repos *Repositories) *cli.Command {
	return &cli.Command{
		Name:  "add",
		Usage: "Add a new track",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "title",
				Aliases:  []string{"t"},
				Usage:    "Track title",
				Required: true,
			},
			&cli.IntFlag{
				Name:     "album",
				Aliases:  []string{"a"},
				Usage:    "Album ID",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "duration",
				Aliases:  []string{"d"},
				Usage:    "Duration in seconds or MM:SS format",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			title := c.String("title")
			albumID := c.Int("album")
			durationStr := c.String("duration")

			duration, err := parseDuration(durationStr)
			if err != nil {
				return err
			}

			track := &models.Track{
				Title:    title,
				AlbumID:  albumID,
				Duration: duration,
			}

			id, err := repos.TrackRepo.Create(track)
			if err != nil {
				return err
			}

			fmt.Printf("Track added with ID: %d\n", id)
			return nil
		},
	}
}
