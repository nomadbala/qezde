-- name: GetAllUsers :many
SELECT * FROM user_schema.users;

-- name: GetUserById :one
SELECT * FROM user_schema.users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO user_schema.users (username, password_hash, salt, email)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateUser :one
UPDATE user_schema.users
SET first_name = $2 AND last_name = $3 AND updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;