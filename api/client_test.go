package api

import (
	mockdb "autoGarage/db/mock"
	db "autoGarage/db/sqlc"
	"autoGarage/util"
	"bytes"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetClientAPI(t *testing.T) {
	client := randomClient()

	testCases := []struct {
		name          string
		clientID      int64
		buildSubs     func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:     "OK",
			clientID: client.ID,
			buildSubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetClient(gomock.Any(), gomock.Eq(client.ID)).Times(1).Return(client, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireMatchBodyClient(t, recorder.Body, client)
			},
		},
		{
			name:     "NotFound",
			clientID: client.ID,
			buildSubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetClient(gomock.Any(), gomock.Eq(client.ID)).Times(1).Return(db.Client{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:     "InternalError",
			clientID: client.ID,
			buildSubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetClient(gomock.Any(), gomock.Eq(client.ID)).Times(1).Return(db.Client{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:     "BadRequest",
			clientID: 0,
			buildSubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetClient(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockStore(ctrl)
			tc.buildSubs(store)
			server := NewServer(store)
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/client/%d", tc.clientID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})

	}
}

func requireMatchBodyClient(t *testing.T, body *bytes.Buffer, client db.Client) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotClient db.Client
	err = json.Unmarshal(data, &gotClient)
	require.NoError(t, err)
	require.Equal(t, gotClient, client)
}

func randomClient() db.Client {
	return db.Client{
		ID:          util.RandomInt(1, 1000),
		Name:        util.RandomClient(),
		Country:     util.RandomCountry(),
		PhoneNumber: util.RandomPhoneNumber(),
	}
}

// mockdb.Mockstore - its all ours interfaces
func TestCreateClient(t *testing.T) {
	client := randomClient()

	testCases := []struct {
		name          string
		body          gin.H
		buildSubs     func(store *mockdb.MockStore)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name":         client.Name,
				"country":      client.Country,
				"phone_number": client.PhoneNumber,
			},
			buildSubs: func(store *mockdb.MockStore) {
				arg := db.CreateClientParams{
					Name:        client.Name,
					Country:     client.Country,
					PhoneNumber: client.PhoneNumber,
				}
				store.EXPECT().CreateClient(gomock.Any(), gomock.Eq(arg)).Times(1).Return(client, nil)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireMatchBodyClient(t, recorder.Body, client)
			},
		},
		{
			name: "BadRequest",
			body: gin.H{
				"name":         client.Name,
				"country":      "",
				"phone_number": client.PhoneNumber,
			},
			buildSubs: func(store *mockdb.MockStore) {
				arg := db.CreateClientParams{
					Name:        client.Name,
					Country:     "",
					PhoneNumber: client.PhoneNumber,
				}
				store.EXPECT().CreateClient(gomock.Any(), gomock.Eq(arg)).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalServerError",
			body: gin.H{
				"name":         client.Name,
				"country":      client.PhoneNumber,
				"phone_number": client.PhoneNumber,
			},
			buildSubs: func(store *mockdb.MockStore) {
				store.EXPECT().CreateClient(gomock.Any(), gomock.Any()).Times(1).Return(db.Client{}, sql.ErrConnDone)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildSubs(store)
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)
			url := "/client"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
