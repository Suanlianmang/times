


-- name: CreateUser :one
INSERT INTO users (email, is_staff, name)
    VALUES ($1, $2, $3)
    ON CONFLICT (email) DO NOTHING
    RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
    WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users
    WHERE id = $1;

