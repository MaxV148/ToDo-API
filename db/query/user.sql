-- name: GetUserByName :one
SELECT *
FROM "user"
WHERE username = $1
LIMIT 1;

-- name: CreateAuthor :one
INSERT INTO "user" (username, password)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteUser :exec
DELETE
FROM "user"
WHERE id = $1;