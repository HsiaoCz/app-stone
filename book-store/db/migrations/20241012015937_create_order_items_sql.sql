-- +goose Up
CREATE TABLE IF NOT EXISTS order_items(
    id integer primary key,
    order_item_id text unique not null,
    order_id text not null references orders,
    book_id text not null references books,
    quantity integer not null,
    price float not null,
    created_at datetime not null,
    updated_at datetime not null,
    deleted_at datetime
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS order_items;
-- +goose StatementBegin
-- +goose StatementEnd