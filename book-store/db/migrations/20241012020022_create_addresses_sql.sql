-- +goose Up
CREATE TABLE IF NOT EXISTS addresses(
    id integer primary key,
    address_id text unique not null,
    user_id text not null references users,
    recipient_name text not null,
    address_line text not null,
    city text not null,
    potal_code integer not null,
    phone_number text not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS addresses;
-- +goose StatementBegin
-- +goose StatementEnd