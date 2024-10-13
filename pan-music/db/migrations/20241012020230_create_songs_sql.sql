-- +goose Up
CREATE TABLE IF NOT EXISTS songs();
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS songs;
-- +goose StatementBegin
-- +goose StatementEnd
