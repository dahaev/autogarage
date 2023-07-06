package api

import (
	mockdb "autoGarage/db/mock"
	db "autoGarage/db/sqlc"
	"autoGarage/util"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type eqCreateUserParameterMatcher struct {
	arg      db.CreateUserParams
	password string
}

func (e eqCreateUserParameterMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}
	err := util.CheckPassword(e.password, arg.HashedPassword)
	if err != nil {
		return false
	}
	e.arg.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParameterMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg db.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParameterMatcher{arg, password}
}

func TestCreateUser(t *testing.T) {
	user, password := CreateRandomUser(t)
	testCases := []struct {
		name          string
		body          gin.H
		buildSubs     func(store *mockdb.MockStore)
		checkResponse func(store *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildSubs: func(store *mockdb.MockStore) {

				hashPassword, err := util.HashPassword(password)
				require.NoError(t, err)

				arg := db.CreateUserParams{
					Username:       user.Username,
					HashedPassword: hashPassword,
					FullName:       user.FullName,
					Email:          user.Email,
				}
				store.EXPECT().CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).Times(1).Return(user, nil)
			},
			checkResponse: func(request *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, request.Code)
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
			url := "/users"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(recorder)

		})
	}

}

func CreateRandomUser(t *testing.T) (user db.User, password string) {
	password = util.RandomUserString(10)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
	user = db.User{
		Username:       util.RandomUserString(10),
		HashedPassword: hashedPassword,
		FullName:       util.RandomUserString(7),
		Email:          util.RandomEmailUser(),
	}
	return
}
