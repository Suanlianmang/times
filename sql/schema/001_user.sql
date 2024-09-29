

-- +goose Up

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v1(),
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    is_staff BOOLEAN DEFAULT TRUE NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255)
);


CREATE TABLE auth_code (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v1(),
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,
    user_id UUID,
    code VARCHAR(6),
    FOREIGN KEY (user_id) REFERENCES users (id)
);

-- +goose Down

DROP EXTENSION IF EXISTS "uuid-ossp";

DROP TABLE users;
DROP TABLE auth_code;


