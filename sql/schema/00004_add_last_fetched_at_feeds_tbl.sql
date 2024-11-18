-- +goose Up
-- +goose StatementBegin
ALTER TABLE feeds
ADD last_fetched_at TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE feeds
DROP last_fetched_at;
-- +goose StatementEnd
