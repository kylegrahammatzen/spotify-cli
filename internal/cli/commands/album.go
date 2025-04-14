package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"spotify-cli/internal/models"
	"time"
)

const dateFormat = "2006-01-02"

func NewAlbumCommand(repos *Repositories) *cli.Command {
	return &cli.Command{
		Name:  "albums",
		Usage: "Manage albums",
		Subcommands: []*cli.Command{
			ListAlbumsCommand(repos),
			AddAlbumCommand(repos),
		},
	}
}

func ListAlbumsCommand(repos *Repositories) *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "List all albums",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "limit",
				Value: 10,
			},
			&cli.IntFlag{
				Name:  "offset",
				Value: 0,
			},
			&cli.IntFlag{
				Name: "artist",
			},
		},
		Action: func(c *cli.Context) error {
			limit := c.Int("limit")
			offset := c.Int("offset")
			artistID := c.Int("artist")

			var albums []*models.Album
			var err error

			if artistID > 0 {
				albums, err = repos.AlbumRepo.GetByArtist(artistID)
			} else {
				albums, err = repos.AlbumRepo.List(limit, offset)
			}

			if err != nil {
				return err
			}

			fmt.Printf("Found %d albums:\n", len(albums))
			for _, album := range albums {
				fmt.Printf("ID: %d, Title: %s, Artist ID: %d, Released: %s\n",
					album.ID, album.Title, album.ArtistID, album.ReleaseDate.Format(dateFormat))
			}

			return nil
		},
	}
}

func AddAlbumCommand(repos *Repositories) *cli.Command {
	return &cli.Command{
		Name:  "add",
		Usage: "Add a new album",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "title",
				Aliases:  []string{"t"},
				Required: true,
			},
			&cli.IntFlag{
				Name:     "artist",
				Aliases:  []string{"a"},
				Required: true,
			},
			&cli.StringFlag{
				Name: "release",
			},
		},
		Action: func(c *cli.Context) error {
			title := c.String("title")
			artistID := c.Int("artist")
			releaseStr := c.String("release")

			var releaseDate time.Time
			var err error
			if releaseStr != "" {
				releaseDate, err = time.Parse(dateFormat, releaseStr)
				if err != nil {
					return fmt.Errorf("invalid date format: %v", err)
				}
			}

			album := &models.Album{
				Title:       title,
				ArtistID:    artistID,
				ReleaseDate: releaseDate,
			}

			id, err := repos.AlbumRepo.Create(album)
			if err != nil {
				return err
			}

			fmt.Printf("Album added with ID: %d\n", id)
			return nil
		},
	}
}
