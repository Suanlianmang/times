


-- name: CreateUser :one
INSERT INTO users (id, email, is_staff, name)
    VALUES ($1, $2, $3, $4)
    RETURNING *;
