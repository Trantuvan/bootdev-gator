-- +goose Up
-- +goose StatementBegin
CREATE TABLE feeds(
    id UUID DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    url TEXT NOT NULL UNIQUE,
    PRIMARY KEY (id),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE feeds;
-- +goose StatementEnd
