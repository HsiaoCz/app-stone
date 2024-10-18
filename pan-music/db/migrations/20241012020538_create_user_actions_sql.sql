-- +goose Up
CREATE TABLE IF NOT EXISTS user_actions(
    id integer primary key,
    action_id text unique not null,
    user_id text not null references users,
    action_type text not null,
    song_id text not null references songs,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS user_actions;
-- +goose StatementBegin
-- +goose StatementEnd