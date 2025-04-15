package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"spotify-cli/internal/models"
)

func NewAuthCommand(repos *Repositories) *cli.Command {
	return &cli.Command{
		Name:  "auth",
		Usage: "Authentication commands",
		Subcommands: []*cli.Command{
			LoginCommand(repos),
			RegisterCommand(repos),
		},
	}
}

func LoginCommand(repos *Repositories) *cli.Command {
	return &cli.Command{
		Name:  "login",
		Usage: "Login with username and password",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "username",
				Aliases:  []string{"u"},
				Usage:    "Username",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "password",
				Aliases:  []string{"p"},
				Usage:    "Password",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			username := c.String("username")
			password := c.String("password")

			user, err := repos.UserRepo.GetByUsername(username)
			if err != nil {
				return fmt.Errorf("login failed: %v", err)
			}

			if user.Password != password {
				return fmt.Errorf("invalid credentials")
			}

			fmt.Printf("Login successful. Welcome, %s!\n", username)
			return nil
		},
	}
}

func RegisterCommand(repos *Repositories) *cli.Command {
	return &cli.Command{
		Name:  "register",
		Usage: "Register a new user",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "username",
				Aliases:  []string{"u"},
				Usage:    "Username",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "password",
				Aliases:  []string{"p"},
				Usage:    "Password",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			username := c.String("username")
			password := c.String("password")

			exists, err := repos.UserRepo.UsernameExists(username)
			if err != nil {
				return fmt.Errorf("error checking username: %v", err)
			}
			if exists {
				return fmt.Errorf("username '%s' is already taken", username)
			}

			user := &models.User{
				Username: username,
				Password: password,
			}

			id, err := repos.UserRepo.Create(user)
			if err != nil {
				return fmt.Errorf("registration failed: %v", err)
			}

			fmt.Printf("User registered successfully with ID: %d\n", id)
			return nil
		},
	}
}
