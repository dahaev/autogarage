-- name: CreateMagazine :one
INSERT INTO magazines (
    address
) VALUES (
             $1
         )
    RETURNING *;

-- name: GetMagazine :one
SELECT * FROM magazines
WHERE id = $1 LIMIT 1;

-- name: ListMagazines :many
SELECT * FROM magazines
ORDER BY id
    LIMIT $1
OFFSET $2;

-- name: UpdateMagazine :one
UPDATE magazines
set address = $2
WHERE id=$1
    RETURNING *;



