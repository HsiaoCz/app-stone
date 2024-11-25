-- +goose Up
CREATE TABLE IF NOT EXISTS authers(
    id integer primary key,
    auther_id text unique not null,
    auther_name text not null,
    bio text not null,
    picture text not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS authers;
-- +goose StatementBegin
-- +goose StatementEnd