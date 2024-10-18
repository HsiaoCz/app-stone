-- +goose Up
CREATE TABLE IF NOT EXISTS comments(
    id integer primary key,
    comment_id text unique not null,
    user_id text not null references users,
    song_id text not null references songs,
    content text not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS comments;
-- +goose StatementBegin
-- +goose StatementEnd