package db

import (
	"autoGarage/util"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func createRandomClient(t *testing.T) Client {
	user := createRandomUser(t)
	arg := CreateClientParams{
		Name:        user.Username,
		Country:     util.RandomCountry(),
		PhoneNumber: util.RandomPhoneNumber(),
	}

	client, err := testQueries.CreateClient(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, client)

	require.Equal(t, arg.Name, arg.Name)
	require.Equal(t, arg.Country, client.Country)
	require.Equal(t, arg.PhoneNumber, client.PhoneNumber)

	require.NotZero(t, client.ID)

	return client
}

func TestCreateClient(t *testing.T) {
	createRandomClient(t)
}

func TestGetClient(t *testing.T) {
	client1 := createRandomClient(t)
	client2, err := testQueries.GetClient(context.Background(), client1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, client2)

	require.Equal(t, client1.ID, client2.ID)
	require.Equal(t, client1.Name, client2.Name)
	require.Equal(t, client1.Country, client2.Country)
	require.Equal(t, client1.PhoneNumber, client2.PhoneNumber)

}

//func TestUpdateClient(t *testing.T) {
//	client1 := createRandomClient(t)
//	arg := UpdateClientParams{
//		ID:          client1.ID,
//		Name:        util.RandomClient(),
//		PhoneNumber: util.RandomPhoneNumber(),
//	}
//	client2, err := testQueries.UpdateClient(context.Background(), arg)
//
//	require.NoError(t, err)
//	require.NotEmpty(t, client2)
//
//	require.Equal(t, client1.ID, client2.ID)
//	require.Equal(t, arg.Name, client2.Name)
//	require.Equal(t, arg.PhoneNumber, client2.PhoneNumber)
//}

func TestListClients(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomClient(t)
	}
	arg := ListClientsParams{
		Limit:  5,
		Offset: 5,
	}

	clients, err := testQueries.ListClients(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, clients)

	for _, client := range clients {
		require.NotEmpty(t, client)
	}
}
