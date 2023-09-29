// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: todo.sql

package db

import (
	"context"
)

const deleteToDo = `-- name: DeleteToDo :exec
DELETE
FROM todo
WHERE id = $1
`

func (q *Queries) DeleteToDo(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteToDo, id)
	return err
}

const listToDoForUser = `-- name: ListToDoForUser :many
SELECT id, title, content, done, created_by, category, created_at
FROM todo
WHERE created_by = $1
LIMIT 1
`

func (q *Queries) ListToDoForUser(ctx context.Context, createdBy int64) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listToDoForUser, createdBy)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.Done,
			&i.CreatedBy,
			&i.Category,
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

const updateToDo = `-- name: UpdateToDo :one
UPDATE todo
set title   = $2,
    content = $3
WHERE id = $1
RETURNING id, title, content, done, created_by, category, created_at
`

type UpdateToDoParams struct {
	ID      int64
	Title   string
	Content string
}

func (q *Queries) UpdateToDo(ctx context.Context, arg UpdateToDoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, updateToDo, arg.ID, arg.Title, arg.Content)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.Done,
		&i.CreatedBy,
		&i.Category,
		&i.CreatedAt,
	)
	return i, err
}
