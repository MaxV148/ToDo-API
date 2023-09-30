-- name: GrantUserToToDo :one
INSERT INTO todo_permissions (user_id, todo_id)
VALUES ($1, $1)
RETURNING *;

-- name: RemoveUserFromToDo :exec
DELETE
FROM todo_permissions
WHERE user_id = $1
  AND todo_id = $2;