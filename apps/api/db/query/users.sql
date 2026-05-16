-- name: GetUser :one
SELECT id, created_at, updated_at
FROM users
WHERE id = $1;
