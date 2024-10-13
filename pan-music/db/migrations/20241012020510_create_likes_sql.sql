-- +goose Up
CREATE TABLE IF NOT EXISTS likes();
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS likes;
-- +goose StatementBegin
-- +goose StatementEnd
