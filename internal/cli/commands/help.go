package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func NewHelpCommand() *cli.Command {
	return &cli.Command{
		Name:   "help",
		Usage:  "Show help information",
		Action: HelpCommand,
	}
}

func HelpCommand(c *cli.Context) error {
	fmt.Println("Spotify CLI - Made by kyle graham matzen")
	fmt.Println("=======================================")
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Println("  spotify-cli [command] [options]")
	fmt.Println()
	fmt.Println("COMMANDS:")
	fmt.Println("  auth       - User authentication (login, register)")
	fmt.Println("  artists    - Manage artists")
	fmt.Println("  albums     - Manage albums")
	fmt.Println("  tracks     - Manage tracks")
	fmt.Println("  playlists  - Manage playlists")
	fmt.Println("  help       - Shows this command lsit")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Println("  spotify-cli auth register --username john --password secret")
	fmt.Println("  spotify-cli artists list")
	fmt.Println("  spotify-cli albums list --artist 1")
	fmt.Println("  spotify-cli tracks add --title \"My Song\" --album 1 --duration 3:45")
	fmt.Println("  spotify-cli playlists create --title \"Favorites\" --user 1")
	fmt.Println()
	fmt.Println("For more information about a command, run:")
	fmt.Println("  spotify-cli [command] --help")

	return nil
}
