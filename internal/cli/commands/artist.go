package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"spotify-cli/internal/models"
)

func NewArtistCommand(repos *Repositories) *cli.Command {
	return &cli.Command{
		Name:  "artists",
		Usage: "Manage artists",
		Subcommands: []*cli.Command{
			ListArtistsCommand(repos),
			AddArtistCommand(repos),
			GetArtistCommand(repos),
		},
	}
}

func ListArtistsCommand(repos *Repositories) *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "List all artists",
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
		},
		Action: func(c *cli.Context) error {
			limit := c.Int("limit")
			offset := c.Int("offset")

			artists, err := repos.ArtistRepo.List(limit, offset)
			if err != nil {
				return err
			}

			fmt.Printf("Found %d artists:\n", len(artists))
			for _, artist := range artists {
				fmt.Printf("ID: %d, Name: %s, Genre: %s\n",
					artist.ID, artist.Name, artist.Genre)
			}

			return nil
		},
	}
}

func AddArtistCommand(repos *Repositories) *cli.Command {
	return &cli.Command{
		Name:  "add",
		Usage: "Add a new artist",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Aliases:  []string{"n"},
				Usage:    "Artist name",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "genre",
				Aliases: []string{"g"},
				Usage:   "Artist genre",
			},
		},
		Action: func(c *cli.Context) error {
			name := c.String("name")
			genre := c.String("genre")

			artist := &models.Artist{
				Name:  name,
				Genre: genre,
			}

			id, err := repos.ArtistRepo.Create(artist)
			if err != nil {
				return err
			}

			fmt.Printf("Artist added with ID: %d\n", id)
			return nil
		},
	}
}

func GetArtistCommand(repos *Repositories) *cli.Command {
	return &cli.Command{
		Name:  "get",
		Usage: "Get artist details",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "id",
				Usage:    "Artist ID",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			id := c.Int("id")

			artist, err := repos.ArtistRepo.GetByID(id)
			if err != nil {
				return err
			}

			fmt.Printf("Artist Details:\n")
			fmt.Printf("ID: %d\n", artist.ID)
			fmt.Printf("Name: %s\n", artist.Name)
			fmt.Printf("Genre: %s\n", artist.Genre)

			return nil
		},
	}
}
