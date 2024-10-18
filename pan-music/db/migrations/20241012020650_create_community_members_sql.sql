-- +goose Up
CREATE TABLE IF NOT EXISTS community_members(
    id integer primary key,
    member_id text unique not null,
    community_id text not null references communities,
    user_id text not null references users,
    joined_at datetime not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS community_members;
-- +goose StatementBegin
-- +goose StatementEnd