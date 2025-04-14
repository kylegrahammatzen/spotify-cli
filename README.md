# Spotify CLI Application

A command-line interface for managing a Spotify-like music streaming service database using MySQL. This project was developed to fulfill the requirements for the Database Systems course project.

# Features

- User management (register, login)
- Music catalog browsing
- Artist, album, and track management
- Playlist creation and manipulation
- Music library organization

## Installation

1. Clone the repository
2. Set up MySQL database using the provided schema file
3. Build the application:

   For Windows:
   ```
   go build -o spotify-cli.exe ./cmd/cli
   ```

   For Mac:
   ```
   go build -o spotify-cli ./cmd/cli
   ```

4. Run the application:

   For Windows:
   ```
   ./spotify-cli.exe [command]
   ```
   For Mac:
   ```
   ./spotify-cli [command]
   ```

## Usage

Run the CLI application using the commands below:

Available commands:
- `auth` - User authentication (login, register)
- `artists` - Manage artists
- `albums` - Manage albums
- `tracks` - Manage tracks
- `playlists` - Manage playlists
- `help` - Show help information