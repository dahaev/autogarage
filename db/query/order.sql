-- name: CreateOrder :one
INSERT INTO orders (
    car,
    client,
    manager,
    magazine,
    delivery_time,
    car_price,
    delivery_price,
    tax,
    total_price
) VALUES (
             $1,$2,$3,$4,$5,$6,$7,$8,$9
         )
    RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1;

-- name: ListOrders :many
SELECT * FROM orders
ORDER BY car
    LIMIT $1
OFFSET $2;

-- name: UpdateOrder :one
UPDATE orders
set car = $2,
    client = $3,
    manager = $4,
    magazine = $5,
    delivery_time = $6,
    car_price = $7,
    delivery_price = $8,
    tax = $9,
    total_price = $10
WHERE id=$1
    RETURNING *;



