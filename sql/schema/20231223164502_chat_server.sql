-- +goose Up
CREATE TABLE IF NOT EXISTS
    chats (
        id SERIAL PRIMARY KEY,
        username VARCHAR(255) NOT NULL,
        msg TEXT NOT NULL,
        created_at TIMESTAMP NOT NULL DEFAULT NOW()
    );

-- +goose Down
DROP TABLE IF EXISTS chats;