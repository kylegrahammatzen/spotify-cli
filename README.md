# Spotify CLI Application

A command-line interface for managing a Spotify-like music streaming service database using MySQL. The project was developed to fulfill the requirements for the Database Systems course project at Rowan University taught by Dominic Boccaleri.

# Features

- User management (register, login)
- Music catalog browsing
- Artist, album, and track management
- Playlist creation and manipulation
- Music library organization

## Installation

1. Clone the repository
2. Set up MySQL database using the provided schema file
3. Create a `.env` file in the project's root directory to store your database credentials. An example `.env` might look like this:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=spotify_cli
```
4. Build the application:

For Windows:
```
go build -o spotify-cli.exe ./cmd/cli
```

For Mac:
```
go build -o spotify-cli ./cmd/cli
```

5. Run the application:

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