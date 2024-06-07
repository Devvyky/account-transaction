-- name: CreateAccount :one
INSERT INTO accounts (
  document_number
) VALUES (
  $1
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE account_id = $1 LIMIT 1;

