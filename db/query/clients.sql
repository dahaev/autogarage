-- name: CreateClient :one
INSERT INTO clients (
    name,
    country,
    phone_number
) VALUES (
             $1, $2, $3
         )
    RETURNING *;

-- name: GetClient :one
SELECT * FROM clients
WHERE id = $1 LIMIT 1;

-- name: ListClients :many
SELECT * FROM clients
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateClient :one
UPDATE clients
set name = $2, phone_number=$3
WHERE id=$1
RETURNING *;



