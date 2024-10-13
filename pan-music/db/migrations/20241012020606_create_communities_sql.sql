-- +goose Up
CREATE TABLE IF NOT EXISTS communities();
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS communities;
-- +goose StatementBegin
-- +goose StatementEnd
