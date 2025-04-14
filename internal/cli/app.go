package cli

import (
	"github.com/urfave/cli/v2"
	"spotify-cli/internal/cli/commands"
	"spotify-cli/internal/repository"
)

type App struct {
	app          *cli.App
	repositories *commands.Repositories
}

func NewApp(
	userRepo *repository.UserRepository,
	artistRepo *repository.ArtistRepository,
	albumRepo *repository.AlbumRepository,
	trackRepo *repository.TrackRepository,
	playlistRepo *repository.PlaylistRepository,
) *App {
	repos := &commands.Repositories{
		UserRepo:     userRepo,
		ArtistRepo:   artistRepo,
		AlbumRepo:    albumRepo,
		TrackRepo:    trackRepo,
		PlaylistRepo: playlistRepo,
	}

	cliApp := &cli.App{
		Name:  "spotify-cli",
		Usage: "Spotify CLI application",
		Commands: []*cli.Command{
			commands.NewAuthCommand(repos),
			commands.NewArtistCommand(repos),
			commands.NewAlbumCommand(repos),
			commands.NewTrackCommand(repos),
			commands.NewPlaylistCommand(repos),
			commands.NewHelpCommand(),
		},
		Action: commands.HelpCommand,
	}

	return &App{
		app:          cliApp,
		repositories: repos,
	}
}

func (a *App) Run(args []string) error {
	return a.app.Run(args)
}
