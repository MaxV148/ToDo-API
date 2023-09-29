// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: user.sql

package db

import (
	"context"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO "user" (username, password)
VALUES ($1, $2)
RETURNING id, username, password, created_at
`

type CreateAuthorParams struct {
	Username string
	Password string
}

func (q *Queries) CreateAuthor(ctx context.Context, arg CreateAuthorParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, arg.Username, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE
FROM "user"
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUserByName = `-- name: GetUserByName :one
SELECT id, username, password, created_at
FROM "user"
WHERE username = $1
LIMIT 1
`

func (q *Queries) GetUserByName(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByName, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}
