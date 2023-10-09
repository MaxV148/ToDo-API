-- name: CreateToDo :one
INSERT INTO todo (title, content, created_by, category)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: ListToDoForUser :many
SELECT todo.id,
       title,
       content,
       done,
       "name" as CategoryName,
       todo.created_at,
       created_by,
       username
FROM todo
         LEFT JOIN todo_permissions ON todo.id = todo_permissions.todo_id AND todo_permissions.user_id = $1
         JOIN category ON todo.category = category.id
         JOIN "user" ON todo.created_by = "user".id
WHERE created_by = $1
   OR todo_permissions.user_id = $1
ORDER BY
    CASE WHEN sqlc.arg(sorting_order) = 'TITLE_ASC' THEN title END,
    CASE WHEN sqlc.arg(sorting_order) = 'TITLE_DESC' THEN title END DESC,
    CASE WHEN sqlc.arg(sorting_order) = 'DONE_ASC' THEN done END,
    CASE WHEN sqlc.arg(sorting_order) = 'DONE_DESC' THEN done END DESC,
    CASE WHEN sqlc.arg(sorting_order) = 'CATEGORY_ASC' THEN name END,
    CASE WHEN sqlc.arg(sorting_order) = 'CATEGORY_DESC' THEN name END DESC,
    CASE WHEN sqlc.arg(sorting_order) = 'CREATED_AT_ASC' THEN todo.created_at END,
    CASE WHEN sqlc.arg(sorting_order) = 'CREATED_AT_DESC' THEN todo.created_at END DESC,
    CASE WHEN sqlc.arg(sorting_order) = 'AUTHOR_ASC' THEN username END,
    CASE WHEN sqlc.arg(sorting_order) = 'AUTHOR_DESC' THEN username END DESC;

-- name: DeleteToDo :one
DELETE
FROM todo
WHERE id = $1
  AND created_by = $2
RETURNING *;

-- name: UpdateToDo :one
UPDATE todo
SET title   = coalesce($2, title),
    content = coalesce($3, content)
WHERE id = $1
RETURNING *;

-- name: ToggleToDoDone :one
UPDATE todo
set done = NOT todo.done
where id = $1
RETURNING *;
