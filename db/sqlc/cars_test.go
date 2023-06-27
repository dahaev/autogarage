package db

import (
	"autoGarage/util"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomCar(t *testing.T) Car {
	arg := CreateCarParams{
		ModelName: util.RandomCar(),
		Equipment: util.RandomEquipment(),
		Color:     util.RandomColor(),
		Country:   util.RandomCountry(),
		Price:     util.RandomPrice(),
		Valid:     "ready",
	}

	car, err := testQueries.CreateCar(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, car)
	require.NotZero(t, car.CreatedAt)

	require.Equal(t, arg.ModelName, car.ModelName)
	require.Equal(t, arg.Equipment, car.Equipment)
	require.Equal(t, arg.Color, car.Color)
	require.Equal(t, arg.Country, car.Country)
	require.Equal(t, arg.Price, car.Price)
	require.Equal(t, arg.Valid, car.Valid)

	return car
}

func TestCreateCar(t *testing.T) {
	createRandomCar(t)
}

func TestGetCar(t *testing.T) {
	car1 := createRandomCar(t)
	car2, err := testQueries.GetCar(context.Background(), car1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, car2)

	require.Equal(t, car1.Valid, car2.Valid)
	require.Equal(t, car1.ModelName, car2.ModelName)
	require.Equal(t, car1.Price, car2.Price)
	require.Equal(t, car1.Equipment, car2.Equipment)
	require.Equal(t, car1.Country, car2.Country)
	require.Equal(t, car1.CreatedAt, car2.CreatedAt)
	require.Equal(t, car1.Color, car2.Color)
}

func TestUpdateCar(t *testing.T) {
	car1 := createRandomCar(t)
	arg := UpdateCarParams{
		ID:        car1.ID,
		Country:   util.RandomCountry(),
		Price:     util.RandomPrice(),
		Color:     util.RandomColor(),
		Valid:     "sold",
		Equipment: util.RandomEquipment(),
	}

	car2, err := testQueries.UpdateCar(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, car2)

	require.Equal(t, arg.Country, car2.Country)
	require.Equal(t, arg.Price, car2.Price)
	require.Equal(t, arg.Color, car2.Color)
	require.Equal(t, arg.Equipment, arg.Equipment)
	require.Equal(t, car1.CreatedAt, car2.CreatedAt)

}

func TestListCars(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCar(t)
	}
	arg := ListCarsParams{
		Limit:  5,
		Offset: 5,
	}
	cars, err := testQueries.ListCars(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cars)

	for _, car := range cars {
		require.NotEmpty(t, car)
	}
}
