-- +goose Up
CREATE TABLE IF NOT EXISTS reviews(
    id integer primary key,
    review_id text unique not null,
    user_id text not null references users,
    book_id text not null references books,
    rating float not null,
    comment text not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS reviews;
-- +goose StatementBegin
-- +goose StatementEnd