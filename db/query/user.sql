-- name: GetUserByName :one
SELECT *
FROM "user"
WHERE username = $1
LIMIT 1;

-- name: CreateUser :one
INSERT INTO "user" (username, password)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteUser :exec
DELETE
FROM "user"
WHERE id = $1;