package token

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var (
	ErrExpiredTokenPaseto = "this token has expired"
)

func TestPasetoMaker(t *testing.T) {
	maker := NewPasetoMaker()

	userId := "cf36766a-3e18-4020-9817-e7103a6ea873"
	duration := time.Minute * 5

	issuedAt := time.Now()
	expiredAt := time.Now().Add(duration)

	token, errCreate := maker.CreateToken(userId, duration)
	require.NoError(t, errCreate)
	require.NotEmpty(t, token)

	payload, errVerify := maker.VerifyToken(token)
	require.NoError(t, errVerify)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, userId, payload.UserID)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	maker := NewPasetoMaker()
	userId := "cf36766a-3e18-4020-9817-e7103a6ea873"

	token, err := maker.CreateToken(userId, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredTokenPaseto)
	require.Nil(t, payload)
}
