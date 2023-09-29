-- name: CreateToDo :one
INSERT INTO todo (title, content, created_by, category)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListToDoForUser :many
SELECT *
FROM todo
WHERE created_by = $1;

-- name: DeleteToDo :exec
DELETE
FROM todo
WHERE id = $1;

-- name: UpdateToDo :one
UPDATE todo
set title   = $2,
    content = $3
WHERE id = $1
RETURNING *;