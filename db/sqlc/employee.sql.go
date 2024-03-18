// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: employee.sql

package db

import (
	"context"
)

const createEmployee = `-- name: CreateEmployee :one
INSERT INTO employee (
  shop_id, user_id, role_id
) VALUES (
  $1, $2, $3
)
RETURNING emp_id, shop_id, user_id, role_id, created_at
`

type CreateEmployeeParams struct {
	ShopID int64
	UserID int64
	RoleID int64
}

func (q *Queries) CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (Employee, error) {
	row := q.db.QueryRowContext(ctx, createEmployee, arg.ShopID, arg.UserID, arg.RoleID)
	var i Employee
	err := row.Scan(
		&i.EmpID,
		&i.ShopID,
		&i.UserID,
		&i.RoleID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEmployee = `-- name: DeleteEmployee :exec
DELETE FROM employee
WHERE emp_id = $1
`

func (q *Queries) DeleteEmployee(ctx context.Context, empID int64) error {
	_, err := q.db.ExecContext(ctx, deleteEmployee, empID)
	return err
}

const getEmployeeByID = `-- name: GetEmployeeByID :one
SELECT emp_id, shop_id, user_id, role_id, created_at FROM employee
WHERE emp_id = $1 LIMIT 1
`

func (q *Queries) GetEmployeeByID(ctx context.Context, empID int64) (Employee, error) {
	row := q.db.QueryRowContext(ctx, getEmployeeByID, empID)
	var i Employee
	err := row.Scan(
		&i.EmpID,
		&i.ShopID,
		&i.UserID,
		&i.RoleID,
		&i.CreatedAt,
	)
	return i, err
}

const listEmployees = `-- name: ListEmployees :many
SELECT emp_id, shop_id, user_id, role_id, created_at FROM employee
`

func (q *Queries) ListEmployees(ctx context.Context) ([]Employee, error) {
	rows, err := q.db.QueryContext(ctx, listEmployees)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Employee
	for rows.Next() {
		var i Employee
		if err := rows.Scan(
			&i.EmpID,
			&i.ShopID,
			&i.UserID,
			&i.RoleID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEmployee = `-- name: UpdateEmployee :one
UPDATE employee
SET 
  shop_id = $2,
  user_id = $3,
  role_id = $4
WHERE emp_id = $1
RETURNING emp_id, shop_id, user_id, role_id, created_at
`

type UpdateEmployeeParams struct {
	EmpID  int64
	ShopID int64
	UserID int64
	RoleID int64
}

func (q *Queries) UpdateEmployee(ctx context.Context, arg UpdateEmployeeParams) (Employee, error) {
	row := q.db.QueryRowContext(ctx, updateEmployee,
		arg.EmpID,
		arg.ShopID,
		arg.UserID,
		arg.RoleID,
	)
	var i Employee
	err := row.Scan(
		&i.EmpID,
		&i.ShopID,
		&i.UserID,
		&i.RoleID,
		&i.CreatedAt,
	)
	return i, err
}