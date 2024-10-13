-- +goose Up
CREATE TABLE IF NOT EXISTS users();
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS users;
-- +goose StatementBegin
-- +goose StatementEnd
