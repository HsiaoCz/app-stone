-- +goose Up
CREATE TABLE IF NOT EXISTS sessions(
    id integer primary key,
    token text unique not null,
    user_id text not null,
    ip_address text not null,
    user_agent text not null,
    expires_at datetime not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS sessions;
-- +goose StatementBegin
-- +goose StatementEnd