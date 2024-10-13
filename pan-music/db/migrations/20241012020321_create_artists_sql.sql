-- +goose Up
CREATE TABLE IF NOT EXISTS artists();
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS artists;
-- +goose StatementBegin
-- +goose StatementEnd
