package db

import (
	"autoGarage/util"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomMagazine(t *testing.T) Magazine {
	arg := util.RandomAddress()
	magazine, err := testQueries.CreateMagazine(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, magazine)

	require.Equal(t, arg, magazine.Address)

	return magazine
}
func TestCreateMagazine(t *testing.T) {
	createRandomMagazine(t)
}

func TestGetMagazine(t *testing.T) {
	magazine1 := createRandomMagazine(t)
	magazine2, err := testQueries.GetMagazine(context.Background(), magazine1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, magazine2)

	require.Equal(t, magazine1.Address, magazine2.Address)
}

func TestUpdateMagazine(t *testing.T) {
	magazine1 := createRandomMagazine(t)
	arg := UpdateMagazineParams{
		ID:      magazine1.ID,
		Address: util.RandomAddress(),
	}
	magazine2, err := testQueries.UpdateMagazine(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, magazine2)

	require.Equal(t, arg.Address, magazine2.Address)
}

func TestListMagazines(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomMagazine(t)
	}
	arg := ListMagazinesParams{
		Limit:  5,
		Offset: 5,
	}

	magazines, err := testQueries.ListMagazines(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, magazines)

	for _, magazine := range magazines {
		require.NotEmpty(t, magazine)
	}
}
