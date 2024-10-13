-- +goose Up
CREATE TABLE IF NOT EXISTS order_items();
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS order_items;
-- +goose StatementBegin
-- +goose StatementEnd