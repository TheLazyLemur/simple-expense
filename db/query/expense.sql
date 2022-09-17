-- name: CreateExpense :one
INSERT INTO expenses (owner, amount)
VALUES ($1, $2) 
RETURNING *;

-- name: GetExpense :one
SELECT * FROM expenses
WHERE id = $1 LIMIT 1;

-- name: ListExpense :many
SELECT * FROM expenses
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteExpense :exec
DELETE FROM expenses
WHERE id = $1;
