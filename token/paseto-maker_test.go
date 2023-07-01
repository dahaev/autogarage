package token

import (
	"autoGarage/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomUserString(32))
	require.NoError(t, err)

	username := util.RandomUserString(10)
	duration := time.Minute
	issuedAt := time.Now()
	expireadAt := issuedAt.Add(duration)
	token, err := maker.CreateToken(username, duration)

	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)

	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAd, time.Second)
	require.WithinDuration(t, expireadAt, payload.ExpiredAt, time.Second)
}
