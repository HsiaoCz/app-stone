-- +goose Up
CREATE TABLE IF NOT EXISTS topics(
    id integer primary key,
    topic_id text unique not null,
    community_id text not null references communities,
    title text not null,
    content text not null,
    user_id text not null references users,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS topics;
-- +goose StatementBegin
-- +goose StatementEnd