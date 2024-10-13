-- +goose Up
CREATE TABLE IF NOT EXISTS user_actions();
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS user_actions;
-- +goose StatementBegin
-- +goose StatementEnd
