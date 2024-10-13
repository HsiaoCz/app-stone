-- +goose Up
CREATE TABLE IF NOT EXISTS topics();
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS topics;
-- +goose StatementBegin
-- +goose StatementEnd