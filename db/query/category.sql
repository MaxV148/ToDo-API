-- name: CreateCategory :one
INSERT INTO category (name, "user")
VALUES ($1, $2)
RETURNING *;

-- name: DeleteCategory :exec
DELETE
FROM category
WHERE id = $1;

-- name: UpdateCategory :one
UPDATE category
set name   = $2
WHERE id = $1
RETURNING *;


-- name: ListCategoriesForUser :many
SELECT *
FROM category
WHERE "user" = $1;