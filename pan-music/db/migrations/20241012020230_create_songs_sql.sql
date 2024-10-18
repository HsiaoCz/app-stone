-- +goose Up
CREATE TABLE IF NOT EXISTS songs(
    id integer primary key,
    song_id text unique not null,
    title text not null,
    artist_id text not null references artists,
    album_id text not null references albums,
    genre text not null,
    durations integer not null,
    release_date text not null,
    lyrics text not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS songs;
-- +goose StatementBegin
-- +goose StatementEnd