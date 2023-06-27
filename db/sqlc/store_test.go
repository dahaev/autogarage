package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateOrderTx(t *testing.T) {
	store := NewStore(testDB)
	car := createRandomCar(t)
	manager := createRandomManager(t)
	magazine := createRandomMagazine(t)
	client := createRandomClient(t)

	n := 2
	results := make(chan OrderTxResult)
	errs := make(chan error)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.CreateOrderTx(context.Background(), CreateOrderTxParams{
				CarID:         car.ID,
				ClientID:      client.ID,
				MagazineID:    magazine.ID,
				ManagerID:     manager.ID,
				CarPrice:      car.Price,
				DeliveryPrice: 20000,
				DeliveryTime:  20,
				Tax:           10000,
				TotalPrice:    car.Price,
			})
			results <- result
			errs <- err
		}()
	}

	for i := 0; i < n; i++ {
		result := <-results
		err := <-errs
		if i == 0 {
			require.NotEmpty(t, result)
			require.NoError(t, err)
		}

		if i == 1 {
			require.NoError(t, err)
			require.Empty(t, result)
		}
	}
}
