package main

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	"spotify-cli/internal/cli"
	"spotify-cli/internal/db"
	"spotify-cli/internal/repository"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	dbConfig := db.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnvAsInt("DB_PORT", 3306),
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", "password"),
		DBName:   getEnv("DB_NAME", "spotify_cli"),
	}

	// Adjust scriptPath to point to the project root, not the cmd/cli folder.
	scriptPath := filepath.Join(getProjectRoot(), "scripts", "db_setup.sql")

	client, err := db.NewMySQLClient(dbConfig, scriptPath)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	defer func() {
		if err := client.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	dbConn := client.GetDB()

	userRepo := repository.NewUserRepository(dbConn)
	artistRepo := repository.NewArtistRepository(dbConn)
	albumRepo := repository.NewAlbumRepository(dbConn)
	trackRepo := repository.NewTrackRepository(dbConn)
	playlistRepo := repository.NewPlaylistRepository(dbConn)

	app := cli.NewApp(
		userRepo,
		artistRepo,
		albumRepo,
		trackRepo,
		playlistRepo,
	)

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Error running application: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}

	return value
}

// getProjectRoot attempts to return the project root directory.
// If you run from "cmd/cli", the project root is assumed to be two directories up.
func getProjectRoot() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}

	// Check if we're running from "cmd/cli".
	// For example, if cwd is "D:\Projects\spotify-cli\cmd\cli"
	// then filepath.Base(cwd) is "cli" and filepath.Base(filepath.Dir(cwd)) is "cmd".
	if filepath.Base(cwd) == "cli" && filepath.Base(filepath.Dir(cwd)) == "cmd" {
		projectRoot, err := filepath.Abs(filepath.Join(cwd, "..", ".."))
		if err != nil {
			log.Fatalf("Failed to determine project root: %v", err)
		}
		return projectRoot
	}

	// Otherwise, return the current working directory.
	return cwd
}
