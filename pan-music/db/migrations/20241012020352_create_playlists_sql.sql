-- +goose Up
CREATE TABLE IF NOT EXISTS playlists();
-- +goose StatementBegin
-- +goose StatementEnd
-- +goose Down
DROP TABLE IF EXISTS playlists;
-- +goose StatementBegin
-- +goose StatementEnd
