CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE artists (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    genre VARCHAR(255),
    image_url VARCHAR(255)
);

CREATE TABLE albums (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    artist_id INT NOT NULL,
    release_date DATE,
    FOREIGN KEY (artist_id) REFERENCES artists(id) ON DELETE CASCADE
);

CREATE TABLE tracks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    album_id INT NOT NULL,
    duration INT NOT NULL,
    FOREIGN KEY (album_id) REFERENCES albums(id) ON DELETE CASCADE
);

CREATE TABLE playlists (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE items (
    playlist_id INT NOT NULL,
    track_id INT NOT NULL,
    position INT NOT NULL,
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (playlist_id, track_id),
    FOREIGN KEY (playlist_id) REFERENCES playlists(id) ON DELETE CASCADE,
    FOREIGN KEY (track_id) REFERENCES tracks(id) ON DELETE CASCADE
);

CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_artists_name ON artists(name);
CREATE INDEX idx_albums_title ON albums(title);
CREATE INDEX idx_albums_artist_id ON albums(artist_id);
CREATE INDEX idx_tracks_title ON tracks(title);
CREATE INDEX idx_tracks_album_id ON tracks(album_id);
CREATE INDEX idx_playlists_title ON playlists(title);
CREATE INDEX idx_playlists_user_id ON playlists(user_id);
CREATE INDEX idx_items_track_id ON items(track_id);

CREATE PROCEDURE CreatePlaylistWithArtistTracks(
    IN p_user_id INT,
    IN p_playlist_title VARCHAR(255),
    IN p_artist_id INT
)
BEGIN
    DECLARE new_playlist_id INT;

    -- Create a new playlist
INSERT INTO playlists (title, user_id)
VALUES (p_playlist_title, p_user_id);

-- Get the id of the newly created playlist
SET new_playlist_id = LAST_INSERT_ID();

    -- Insert tracks into the playlist (all tracks from albums by the artist)
INSERT INTO items (playlist_id, track_id, position)
SELECT
    new_playlist_id,
    t.id AS track_id,
    ROW_NUMBER() OVER (ORDER BY t.id) AS position
FROM tracks t
    JOIN albums a ON t.album_id = a.id
WHERE a.artist_id = p_artist_id;
END;

-- CALL CreatePlaylistWithArtistTracks(1, 'My Favorite Artist', 5);