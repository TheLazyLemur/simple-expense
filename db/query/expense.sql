-- name: CreateExpense :one
INSERT INTO expenses (organisation_id, uploader, amount)
VALUES ($1, $2, $3) 
RETURNING *;

-- name: GetExpense :one
SELECT * FROM expenses
WHERE id = $1 LIMIT 1;

-- name: ListExpenseByOrganisation :many
SELECT * FROM expenses
WHERE organisation_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: ListExpense :many
SELECT * FROM expenses
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteExpense :exec
DELETE FROM expenses
WHERE id = $1;
