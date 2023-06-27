package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

//go:generate mockgen -source=store.go -destination=mocks/mock.go
type Store interface {
	Querier
	CreateOrderTx(ctx context.Context, arg CreateOrderTxParams) (OrderTxResult, error)
}
type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type CreateOrderTxParams struct {
	CarID         int64 `json:"car_id"`
	ClientID      int64 `json:"client_id"`
	ManagerID     int64 `json:"manager_id"`
	MagazineID    int64 `json:"magazine_id"`
	DeliveryTime  int32 `json:"delivery_time"`
	CarPrice      int64 `json:"car_price"`
	DeliveryPrice int64 `json:"delivery_price"`
	Tax           int64 `json:"tax"`
	TotalPrice    int64 `json:"total_price"`
}

type OrderTxResult struct {
	Order Order `json:"order"`
	Car   Car   `json:"car"`
}

func (store *SQLStore) CreateOrderTx(ctx context.Context, arg CreateOrderTxParams) (OrderTxResult, error) {
	var result OrderTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		car, err := q.GetCarForUpdate(ctx, arg.CarID)
		if err != nil {
			return err
		}
		if car.Valid == "sold" {
			log.Println("Car get allready")
			return err
		}

		result.Order, err = q.CreateOrder(ctx, CreateOrderParams{
			Car:           arg.CarID,
			Client:        arg.ClientID,
			Manager:       arg.ManagerID,
			Magazine:      arg.MagazineID,
			DeliveryTime:  arg.DeliveryTime,
			DeliveryPrice: arg.DeliveryPrice,
			CarPrice:      arg.CarPrice,
			Tax:           arg.Tax,
			TotalPrice:    arg.TotalPrice,
		})
		if err != nil {
			return err
		}
		result.Car, err = q.UpdateCar(ctx, UpdateCarParams{
			ID:        arg.CarID,
			ModelName: car.ModelName,
			Equipment: car.Equipment,
			Color:     car.Color,
			Country:   car.Country,
			Price:     car.Price,
			Valid:     "sold",
		})
		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}
