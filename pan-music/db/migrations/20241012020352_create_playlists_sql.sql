-- +goose Up
CREATE TABLE IF NOT EXISTS playlists(
    id integer primary key,
    playlist_id text unique not null,
    user_id text not null references users,
    title text not null,
    descriptions text not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS playlists;
-- +goose StatementBegin
-- +goose StatementEnd