-- +goose Up
CREATE TABLE IF NOT EXISTS community_members();
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS community_members;
-- +goose StatementBegin
-- +goose StatementEnd