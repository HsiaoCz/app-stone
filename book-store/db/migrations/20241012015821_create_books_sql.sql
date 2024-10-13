-- +goose Up
CREATE TABLE IF NOT EXISTS books(

);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS books;
-- +goose StatementBegin
-- +goose StatementEnd
