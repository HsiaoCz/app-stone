-- +goose Up
CREATE TABLE IF NOT EXISTS records(
    id integer primary key,
    record_id text unique not null,
    user_id text not null,
    book_id text not null,
    type_name text not null,
    device text not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS records;
-- +goose StatementBegin
-- +goose StatementEnd