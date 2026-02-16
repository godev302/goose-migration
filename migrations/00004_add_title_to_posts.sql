-- +goose Up
-- +goose StatementBegin
ALTER TABLE posts ADD COLUMN title TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE posts DROP COLUMN title;
-- +goose StatementEnd
