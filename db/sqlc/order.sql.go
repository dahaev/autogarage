// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: order.sql

package db

import (
	"context"
)

const createOrder = `-- name: CreateOrder :one
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
    RETURNING id, car, client, manager, magazine, delivery_time, car_price, delivery_price, tax, total_price, created_at
`

type CreateOrderParams struct {
	Car           int64 `json:"car"`
	Client        int64 `json:"client"`
	Manager       int64 `json:"manager"`
	Magazine      int64 `json:"magazine"`
	DeliveryTime  int32 `json:"delivery_time"`
	CarPrice      int64 `json:"car_price"`
	DeliveryPrice int64 `json:"delivery_price"`
	Tax           int64 `json:"tax"`
	TotalPrice    int64 `json:"total_price"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder,
		arg.Car,
		arg.Client,
		arg.Manager,
		arg.Magazine,
		arg.DeliveryTime,
		arg.CarPrice,
		arg.DeliveryPrice,
		arg.Tax,
		arg.TotalPrice,
	)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.Car,
		&i.Client,
		&i.Manager,
		&i.Magazine,
		&i.DeliveryTime,
		&i.CarPrice,
		&i.DeliveryPrice,
		&i.Tax,
		&i.TotalPrice,
		&i.CreatedAt,
	)
	return i, err
}

const getOrder = `-- name: GetOrder :one
SELECT id, car, client, manager, magazine, delivery_time, car_price, delivery_price, tax, total_price, created_at FROM orders
WHERE id = $1
`

func (q *Queries) GetOrder(ctx context.Context, id int64) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrder, id)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.Car,
		&i.Client,
		&i.Manager,
		&i.Magazine,
		&i.DeliveryTime,
		&i.CarPrice,
		&i.DeliveryPrice,
		&i.Tax,
		&i.TotalPrice,
		&i.CreatedAt,
	)
	return i, err
}

const listOrders = `-- name: ListOrders :many
SELECT id, car, client, manager, magazine, delivery_time, car_price, delivery_price, tax, total_price, created_at FROM orders
ORDER BY car
    LIMIT $1
OFFSET $2
`

type ListOrdersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListOrders(ctx context.Context, arg ListOrdersParams) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, listOrders, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Order
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.ID,
			&i.Car,
			&i.Client,
			&i.Manager,
			&i.Magazine,
			&i.DeliveryTime,
			&i.CarPrice,
			&i.DeliveryPrice,
			&i.Tax,
			&i.TotalPrice,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrder = `-- name: UpdateOrder :one
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
    RETURNING id, car, client, manager, magazine, delivery_time, car_price, delivery_price, tax, total_price, created_at
`

type UpdateOrderParams struct {
	ID            int64 `json:"id"`
	Car           int64 `json:"car"`
	Client        int64 `json:"client"`
	Manager       int64 `json:"manager"`
	Magazine      int64 `json:"magazine"`
	DeliveryTime  int32 `json:"delivery_time"`
	CarPrice      int64 `json:"car_price"`
	DeliveryPrice int64 `json:"delivery_price"`
	Tax           int64 `json:"tax"`
	TotalPrice    int64 `json:"total_price"`
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, updateOrder,
		arg.ID,
		arg.Car,
		arg.Client,
		arg.Manager,
		arg.Magazine,
		arg.DeliveryTime,
		arg.CarPrice,
		arg.DeliveryPrice,
		arg.Tax,
		arg.TotalPrice,
	)
	var i Order
	err := row.Scan(
		&i.ID,
		&i.Car,
		&i.Client,
		&i.Manager,
		&i.Magazine,
		&i.DeliveryTime,
		&i.CarPrice,
		&i.DeliveryPrice,
		&i.Tax,
		&i.TotalPrice,
		&i.CreatedAt,
	)
	return i, err
}