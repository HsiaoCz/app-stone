-- +goose Up
CREATE TABLE IF NOT EXISTS orders(

);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS orders;
-- +goose StatementBegin
-- +goose StatementEnd
