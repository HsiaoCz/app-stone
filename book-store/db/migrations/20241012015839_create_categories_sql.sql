-- +goose Up
CREATE TABLE IF NOT EXISTS categories(

);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS categories;
-- +goose StatementBegin
-- +goose StatementEnd
