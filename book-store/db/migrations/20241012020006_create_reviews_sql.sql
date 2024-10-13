-- +goose Up
CREATE TABLE IF NOT EXISTS reviews(

);
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS reviews;
-- +goose StatementBegin
-- +goose StatementEnd
