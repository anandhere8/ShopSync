// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"time"
)

type User struct {
	ID        int64
	Username  string
	Role      string
	CreatedAt time.Time
}