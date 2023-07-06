package api

import (
	mockdb "autoGarage/db/mock"
	db "autoGarage/db/sqlc"
	"autoGarage/token"
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
	"time"
)

func TestGetClientAPI(t *testing.T) {
	user, _ := CreateRandomUser(t)
	client := randomClient(user.Username)

	testCases := []struct {
		name          string
		clientID      int64
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildSubs     func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:     "OK",
			clientID: client.ID,
			buildSubs: func(store *mockdb.MockStore) {
				store.EXPECT().GetClient(gomock.Any(), gomock.Eq(client.ID)).Times(1).Return(client, nil)
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				authorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
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
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				authorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
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
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				authorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
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
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				authorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
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
			server, err := NewServer(store)
			require.NoError(t, err)
			recorder := httptest.NewRecorder()
			url := fmt.Sprintf("/client/%d", tc.clientID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)
			tc.setupAuth(t, request, server.tokenMaker)
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

func randomClient(user string) db.Client {
	return db.Client{
		ID:          util.RandomInt(1, 1000),
		Name:        user,
		Country:     util.RandomCountry(),
		PhoneNumber: util.RandomPhoneNumber(),
	}
}

// mockdb.Mockstore - its all ours interfaces
func TestCreateClientAPI(t *testing.T) {
	user, _ := CreateRandomUser(t)
	client := randomClient(user.Username)

	testCases := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
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
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				authorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
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
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				authorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
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
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				authorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
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
			server, err := NewServer(store)
			require.NoError(t, err)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)
			url := "/client"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)
			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)
		})
	}
}
