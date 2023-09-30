-- name: GrantUserToToDo :one
INSERT INTO todo_permissions (user_id, todo_id)
VALUES ($1, $2)
RETURNING *;

-- name: RemoveUserFromToDo :exec
DELETE
FROM todo_permissions
WHERE user_id = $1
  AND todo_id = $2;

-- name: ListGrantedToDoForUser :many
SELECT *
FROM todo
INNER JOIN todo_permissions ON todo.id = todo_permissions.todo_id AND todo_permissions.user_id = $1;
