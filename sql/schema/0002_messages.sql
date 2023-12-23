-- +goose Up
CREATE TABLE IF NOT EXISTS
    messages (
        id SERIAL PRIMARY KEY,
        username VARCHAR(255) NOT NULL,
        msg TEXT NOT NULL,
        published_at TIMESTAMP NOT NULL DEFAULT NOW(),
        chat_id SERIAL NOT NULL REFERENCES chats (id) ON DELETE CASCADE,
        UNIQUE(username,chat_id)
    );

-- +goose Down
DROP TABLE IF EXISTS messages;