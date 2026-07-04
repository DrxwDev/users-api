-- name: Save :one
INSERT INTO users (id, email, password) VALUES ($1, $2, $3)
RETURNING id, email, created_at, updated_at;

-- name: FindByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: FindByID :one
SELECT * FROM users WHERE id = $1;
