-- +goose Up
CREATE TABLE IF NOT EXISTS comments(
    
);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS comments;
-- +goose StatementBegin
-- +goose StatementEnd
