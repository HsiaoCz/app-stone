-- +goose Up
CREATE TABLE IF NOT EXISTS books(
    id integer primary key,
    book_id text unique not null,
    title text not null,
    auther_id text not null references authers,
    price float not null,
    stock integer not null,
    category_id text not null references categories,
    descriptions text not null,
    cover_image text not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS books;
-- +goose StatementBegin
-- +goose StatementEnd