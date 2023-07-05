-- name: CreateUser :one
INSERT INTO users(user_id, username, password, created_at)
VALUES($1, $2, $3, $4)
RETURNING *;

-- name: GetUserByUsername :one
SELECT
  u.user_id
  ,username
  ,password
  ,created_at
FROM users as u
WHERE u.username = $1;