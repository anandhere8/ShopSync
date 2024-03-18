-- name: GetUserByID :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: CreateUser :one
INSERT INTO users (
  firstname, lastname, username, email, phone_number, password_hash
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET 
  firstname = $2,
  lastname = $3,
  username = $4,
  email = $5,
  phone_number = $6,
  password_hash = $7
WHERE user_id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1;




