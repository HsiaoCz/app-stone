-- +goose Up
CREATE TABLE IF NOT EXISTS users(
    id integer primary key,
    user_id text unique not null,
    username text not null,
    email text unique not null,
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS users;
-- +goose StatementBegin
-- +goose StatementEnd
