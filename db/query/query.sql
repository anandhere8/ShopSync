-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username;

-- name: CreateUser :one
INSERT INTO users (
  username, role
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
  set username = $2,
  role = $3
WHERE id = $1
RETURNING *;


-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;