// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"context"
)

type Querier interface {
	CreateCar(ctx context.Context, arg CreateCarParams) (Car, error)
	CreateClient(ctx context.Context, arg CreateClientParams) (Client, error)
	CreateMagazine(ctx context.Context, address string) (Magazine, error)
	CreateManager(ctx context.Context, arg CreateManagerParams) (Manager, error)
	CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetCar(ctx context.Context, id int64) (Car, error)
	GetCarForUpdate(ctx context.Context, id int64) (Car, error)
	GetClient(ctx context.Context, id int64) (Client, error)
	GetMagazine(ctx context.Context, id int64) (Magazine, error)
	GetManager(ctx context.Context, id int64) (Manager, error)
	GetOrder(ctx context.Context, id int64) (Order, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListCars(ctx context.Context, arg ListCarsParams) ([]Car, error)
	ListClients(ctx context.Context, arg ListClientsParams) ([]Client, error)
	ListMagazines(ctx context.Context, arg ListMagazinesParams) ([]Magazine, error)
	ListManagers(ctx context.Context, arg ListManagersParams) ([]Manager, error)
	ListOrders(ctx context.Context, arg ListOrdersParams) ([]Order, error)
	UpdateCar(ctx context.Context, arg UpdateCarParams) (Car, error)
	UpdateClient(ctx context.Context, arg UpdateClientParams) (Client, error)
	UpdateMagazine(ctx context.Context, arg UpdateMagazineParams) (Magazine, error)
	UpdateManager(ctx context.Context, arg UpdateManagerParams) (Manager, error)
	UpdateOrder(ctx context.Context, arg UpdateOrderParams) (Order, error)
}

var _ Querier = (*Queries)(nil)
