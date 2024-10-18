-- +goose Up
CREATE TABLE IF NOT EXISTS topic_comments(
    id integer primary key,
    comment_id text unique not null,
    topic_id text not null references topics,
    user_id text not null references users,
    content text not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS topic_comments;
-- +goose StatementBegin
-- +goose StatementEnd