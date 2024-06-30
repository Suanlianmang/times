

-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    is_staff BOOLEAN DEFAULT TRUE NOT NULL,
    email VARCHAR(255) NOT NULL,
    name VARCHAR(255)
);

-- +goose Down

DROP TABLE users;


