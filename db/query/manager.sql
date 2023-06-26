-- name: CreateManager :one
INSERT INTO managers (
    name,
    town
) VALUES (
             $1, $2
         )
    RETURNING *;

-- name: GetManager :one
SELECT * FROM managers
WHERE id = $1 LIMIT 1;

-- name: ListManagers :many
SELECT * FROM managers
ORDER BY id
    LIMIT $1
OFFSET $2;

-- name: UpdateManager :one
UPDATE managers
set name = $2, town=$3
WHERE id=$1
    RETURNING *;



