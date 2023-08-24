-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmailOrUsername :one
SELECT * FROM users
WHERE email = $1 OR username = $1
LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
LIMIT $1
OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (
  full_name, username, email, hashed_password 
) VALUES (
  $1, $2, $3, $4 
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
