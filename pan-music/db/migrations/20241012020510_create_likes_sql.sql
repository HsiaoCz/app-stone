-- +goose Up
CREATE TABLE IF NOT EXISTS likes(
    id integer primary key,
    like_id text unique not null,
    user_id text not null references users,
    song_id text not null references songs,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS likes;
-- +goose StatementBegin
-- +goose StatementEnd
