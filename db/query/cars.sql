-- name: CreateCar :one
INSERT INTO cars (
        model_name,
        equipment,
        color,
        country,
        price,
        valid

) VALUES (
             $1,$2,$3,$4,$5,$6
         )
    RETURNING *;

-- name: GetCar :one
SELECT * FROM cars
WHERE id = $1;

-- name: GetCarForUpdate :one
SELECT * FROM cars
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListCars :many
SELECT * FROM cars
ORDER BY model_name
    LIMIT $1
OFFSET $2;

-- name: UpdateCar :one
UPDATE cars
set model_name = $2,
    equipment = $3,
    color = $4,
    country = $5,
    price = $6,
    valid = $7
WHERE id=$1
    RETURNING *;




