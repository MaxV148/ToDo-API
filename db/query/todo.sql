-- name: ListToDoForUser :many
SELECT *
FROM todo
WHERE created_by = $1
LIMIT 1;

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