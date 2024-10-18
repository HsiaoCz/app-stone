-- +goose Up
CREATE TABLE IF NOT EXISTS artists(
    id integer primary key,
    artist_id text unique not null,
    name text not null,
    bio text not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS artists;
-- +goose StatementBegin
-- +goose StatementEnd
