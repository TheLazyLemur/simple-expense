-- name: CreateInvoice :one
INSERT INTO invoices (owner, expense_id, url)
VALUES ($1, $2, $3) 
RETURNING *;

-- name: GetInvoice :one
SELECT * FROM invoices
WHERE id = $1 LIMIT 1;

-- name: ListInvoice :many
SELECT * FROM invoices
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteInvoice :exec
DELETE FROM invoices
WHERE id = $1;
