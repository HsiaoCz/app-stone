-- +goose Up
CREATE TABLE IF NOT EXISTS albums(
    id integer primary key,
    album_id text unique not null,
    title text not null,
    artist_id text not null references artists,
    release_date text not null,
    cover_image text not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS albums;
-- +goose StatementBegin
-- +goose StatementEnd