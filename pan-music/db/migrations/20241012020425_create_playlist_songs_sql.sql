-- +goose Up
CREATE TABLE IF NOT EXISTS playlist_songs();
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS playlist_songs;
-- +goose StatementBegin
-- +goose StatementEnd
