// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: todoPermission.sql

package db

import (
	"context"
	"time"
)

const grantUserToToDo = `-- name: GrantUserToToDo :one
INSERT INTO todo_permissions (user_id, todo_id)
VALUES ($1, $2)
RETURNING user_id, todo_id
`

type GrantUserToToDoParams struct {
	UserID int64
	TodoID int64
}

func (q *Queries) GrantUserToToDo(ctx context.Context, arg GrantUserToToDoParams) (TodoPermission, error) {
	row := q.db.QueryRowContext(ctx, grantUserToToDo, arg.UserID, arg.TodoID)
	var i TodoPermission
	err := row.Scan(&i.UserID, &i.TodoID)
	return i, err
}

const listGrantedToDoForUser = `-- name: ListGrantedToDoForUser :many
SELECT id, title, content, done, created_by, category, created_at, user_id, todo_id
FROM todo
INNER JOIN todo_permissions ON todo.id = todo_permissions.todo_id AND todo_permissions.user_id = $1
`

type ListGrantedToDoForUserRow struct {
	ID        int64
	Title     string
	Content   string
	Done      bool
	CreatedBy int64
	Category  int64
	CreatedAt time.Time
	UserID    int64
	TodoID    int64
}

func (q *Queries) ListGrantedToDoForUser(ctx context.Context, userID int64) ([]ListGrantedToDoForUserRow, error) {
	rows, err := q.db.QueryContext(ctx, listGrantedToDoForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListGrantedToDoForUserRow{}
	for rows.Next() {
		var i ListGrantedToDoForUserRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.Done,
			&i.CreatedBy,
			&i.Category,
			&i.CreatedAt,
			&i.UserID,
			&i.TodoID,
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

const removeUserFromToDo = `-- name: RemoveUserFromToDo :exec
DELETE
FROM todo_permissions
WHERE user_id = $1
  AND todo_id = $2
`

type RemoveUserFromToDoParams struct {
	UserID int64
	TodoID int64
}

func (q *Queries) RemoveUserFromToDo(ctx context.Context, arg RemoveUserFromToDoParams) error {
	_, err := q.db.ExecContext(ctx, removeUserFromToDo, arg.UserID, arg.TodoID)
	return err
}
