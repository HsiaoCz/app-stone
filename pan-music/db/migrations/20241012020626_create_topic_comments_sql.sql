-- +goose Up
CREATE TABLE IF NOT EXISTS topic_comments();
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS topic_comments;
-- +goose StatementBegin
-- +goose StatementEnd