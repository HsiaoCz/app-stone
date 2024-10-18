-- +goose Up
CREATE TABLE IF NOT EXISTS playlist_songs(
    id integer primary key,
    playlist_id text not null references playlists,
    song_id text not null references songs,
    added_at datetime not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS playlist_songs;
-- +goose StatementBegin
-- +goose StatementEnd