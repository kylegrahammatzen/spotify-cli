package main

import (
	"fmt"
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
	fmt.Println("Starting Spotify CLI application...")

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	dbConfig := db.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnvAsInt("DB_PORT", 3306),
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", "password"),
		DBName:   getEnv("DB_NAME", "spotify_cli"),
	}

	client, err := db.NewMySQLClient(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	defer func() {
		if err := client.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()

	scriptPath := filepath.Join("scripts", "db_setup.sql")
	if _, err := os.Stat(scriptPath); err == nil {
		fmt.Println("Setting up database...")
		if err := client.SetupDB(dbConfig.DBName, scriptPath); err != nil {
			log.Fatalf("Failed to set up database: %v", err)
		}
		fmt.Println("Database setup complete")
	}

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
