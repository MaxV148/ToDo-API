// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: category.sql

package db

import (
	"context"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO category (name, "user")
VALUES ($1, $2)
RETURNING id, name, "user", created_at
`

type CreateCategoryParams struct {
	Name string
	User int64
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, createCategory, arg.Name, arg.User)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.User,
		&i.CreatedAt,
	)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE
FROM category
WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const listCategoriesForUser = `-- name: ListCategoriesForUser :many
SELECT id, name, "user", created_at
FROM category
WHERE "user" = $1
`

func (q *Queries) ListCategoriesForUser(ctx context.Context, user int64) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, listCategoriesForUser, user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.User,
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

const updateCategory = `-- name: UpdateCategory :one
UPDATE category
set name   = $2
WHERE id = $1
RETURNING id, name, "user", created_at
`

type UpdateCategoryParams struct {
	ID   int64
	Name string
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, updateCategory, arg.ID, arg.Name)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.User,
		&i.CreatedAt,
	)
	return i, err
}
