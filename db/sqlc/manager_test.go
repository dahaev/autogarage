package db

import (
	"autoGarage/util"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomManager(t *testing.T) Manager {
	arg := CreateManagerParams{
		Name: util.RandomClient(),
		Town: util.RandomCity(),
	}
	manager, err := testQueries.CreateManager(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, manager)

	require.Equal(t, arg.Name, manager.Name)
	require.Equal(t, arg.Town, manager.Town)

	return manager
}

func TestCreateManager(t *testing.T) {
	createRandomManager(t)
}

func TestUpdateManager(t *testing.T) {
	manager1 := createRandomManager(t)
	arg := UpdateManagerParams{
		ID:   manager1.ID,
		Name: util.RandomClient(),
		Town: util.RandomCity(),
	}

	manager2, err := testQueries.UpdateManager(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, manager2)

	require.Equal(t, arg.Name, manager2.Name)
	require.Equal(t, arg.Town, manager2.Town)
}

func TestListManagers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomManager(t)
	}
	arg := ListManagersParams{
		Limit:  5,
		Offset: 5,
	}
	managers, err := testQueries.ListManagers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, managers)

	for _, manager := range managers {
		require.NotEmpty(t, manager)
	}
}
