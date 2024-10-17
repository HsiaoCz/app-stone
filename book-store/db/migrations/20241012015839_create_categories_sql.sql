-- +goose Up
CREATE TABLE IF NOT EXISTS categories(
    id integer primary key,
    category_id integer unique not null,
    category_name text unique not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS categories;
-- +goose StatementBegin
-- +goose StatementEnd