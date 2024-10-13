-- +goose Up
CREATE TABLE IF NOT EXISTS addresses();
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS addresses;
-- +goose StatementBegin
-- +goose StatementEnd