-- +goose Up
CREATE TABLE IF NOT EXISTS albums();
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS albums;
-- +goose StatementBegin
-- +goose StatementEnd
