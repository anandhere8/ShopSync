-- name: GetEmployeeByID :one
SELECT * FROM employee
WHERE emp_id = $1 LIMIT 1;

-- name: ListEmployees :many
SELECT * FROM employee;

-- name: CreateEmployee :one
INSERT INTO employee (
  shop_id, user_id, role_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateEmployee :one
UPDATE employee
SET 
  shop_id = $2,
  user_id = $3,
  role_id = $4
WHERE emp_id = $1
RETURNING *;

-- name: DeleteEmployee :exec
DELETE FROM employee
WHERE emp_id = $1;