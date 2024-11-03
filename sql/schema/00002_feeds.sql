-- +goose Up
-- +goose StatementBegin
CREATE TABLE feeds(
    id UUID DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID REFERENCES users (id) ON DELETE CASCADE,
    PRIMARY KEY (id),
    FOREIGN KEY(user_id) REFERENCES users (id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE feeds;
-- +goose StatementEnd
