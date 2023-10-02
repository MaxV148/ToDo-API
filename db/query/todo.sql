-- name: CreateToDo :one
INSERT INTO todo (title, content, created_by, category)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListToDoForUser :many
SELECT todo.id         as ToDoId,
       title           as ToDoTitle,
       content         as ToDoContent,
       done,
       "name"          as categoryName,
       todo.created_at as ToDoCreatedAt,
       created_by      as ToDoCreatedBy
FROM todo
         JOIN category on todo.category = category.id
WHERE created_by = $1
ORDER BY categoryName DESC;

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

-- name: ToggleToDoDone :one
UPDATE todo
set done = NOT todo.done
where id = $1
RETURNING *;


