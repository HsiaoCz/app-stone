-- +goose Up
CREATE TABLE IF NOT EXISTS orders(
    id integer primary key,
    order_id text unique not null,
    user_id text not null references users,
    order_status text not null,
    total_amount float not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS orders;
-- +goose StatementBegin
-- +goose StatementEnd