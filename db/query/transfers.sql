CREATE TABLE "transfers" (
  "id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "amount" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id, 
  to_account_id, 
  amount
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfersByFromAccount :many
SELECT * FROM transfers
WHERE from_account_id = $1 
LIMIT $2
OFFSET $3;

-- name: ListTransfersByToAccount :many
SELECT * FROM transfers
WHERE to_account_id = $1 
LIMIT $2
OFFSET $3;

-- name: ListTransfersByFromAndToAccount :many
SELECT * FROM transfers
WHERE from_account_id = $1 AND to_account_id = $2 
LIMIT $3
OFFSET $4;

-- name: ListTransfer :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTransfer :one
UPDATE transfers SET amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTransfer :exec
DELETE FROM transfers WHERE id = $1;

-- name: DeleteTransfersByFromAccount :exec
DELETE FROM transfers WHERE from_account_id = $1;

-- name: DeleteTransfersByToAccount :exec
DELETE FROM transfers WHERE to_account_id = $1;

-- name: DeleteTransfersByFromAccountAndToAccount :exec
DELETE FROM transfers WHERE from_account_id = $1 AND to_account_id = $2;