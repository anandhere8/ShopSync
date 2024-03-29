-- name: GetRoleByID :one
SELECT * FROM roles
WHERE role_id = $1 LIMIT 1;

-- name: ListRoles :many
SELECT * FROM roles;

-- name: CreateRole :one
INSERT INTO roles (
  role_name
) VALUES (
  $1
)
RETURNING *;

-- name: UpdateRole :one
UPDATE roles
SET 
  role_name = $2
WHERE role_id = $1
RETURNING *;

-- name: DeleteRole :exec
DELETE FROM roles
WHERE role_id = $1;