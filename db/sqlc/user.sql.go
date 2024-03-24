// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  firstname, lastname, username, email, phone_number, password_hash
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING user_id, firstname, lastname, username, email, phone_number, password_hash, created_at
`

type CreateUserParams struct {
	Firstname    string
	Lastname     string
	Username     string
	Email        string
	PhoneNumber  string
	PasswordHash string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Firstname,
		arg.Lastname,
		arg.Username,
		arg.Email,
		arg.PhoneNumber,
		arg.PasswordHash,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Firstname,
		&i.Lastname,
		&i.Username,
		&i.Email,
		&i.PhoneNumber,
		&i.PasswordHash,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, userID int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, userID)
	return err
}

const getUserByID = `-- name: GetUserByID :one
SELECT user_id, firstname, lastname, username, email, phone_number, password_hash, created_at FROM users
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, userID int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Firstname,
		&i.Lastname,
		&i.Username,
		&i.Email,
		&i.PhoneNumber,
		&i.PasswordHash,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT user_id, firstname, lastname, username, email, phone_number, password_hash, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Firstname,
		&i.Lastname,
		&i.Username,
		&i.Email,
		&i.PhoneNumber,
		&i.PasswordHash,
		&i.CreatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT user_id, firstname, lastname, username, email, phone_number, password_hash, created_at FROM users
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Firstname,
			&i.Lastname,
			&i.Username,
			&i.Email,
			&i.PhoneNumber,
			&i.PasswordHash,
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

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET 
  firstname = $2,
  lastname = $3,
  username = $4,
  email = $5,
  phone_number = $6,
  password_hash = $7
WHERE user_id = $1
RETURNING user_id, firstname, lastname, username, email, phone_number, password_hash, created_at
`

type UpdateUserParams struct {
	UserID       int64
	Firstname    string
	Lastname     string
	Username     string
	Email        string
	PhoneNumber  string
	PasswordHash string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.UserID,
		arg.Firstname,
		arg.Lastname,
		arg.Username,
		arg.Email,
		arg.PhoneNumber,
		arg.PasswordHash,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Firstname,
		&i.Lastname,
		&i.Username,
		&i.Email,
		&i.PhoneNumber,
		&i.PasswordHash,
		&i.CreatedAt,
	)
	return i, err
}
