-- +goose Up
CREATE TABLE IF NOT EXISTS communities(
    id integer primary key,
    community_id text unique not null,
    community_name text  not null,
    descriptions text not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS communities;
-- +goose StatementBegin
-- +goose StatementEnd