-- name: CreateOrganisation :one
INSERT INTO organisations (name, owner)
VALUES ($1, $2) 
RETURNING *;

-- name: GetOrganisation :one
SELECT * FROM organisations
WHERE id = $1 LIMIT 1;

-- name: ListOrganisation :many
SELECT * FROM organisations
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateOrganisation :one
UPDATE organisations
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteOrganisation :exec
DELETE FROM organisations
WHERE id = $1;
