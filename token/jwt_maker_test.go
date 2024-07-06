package token

import (
	"testing"
	"time"

	"github.com/PyMarcus/go_sqlc/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	secretKey := util.RandomString(32)
	maker, err := NewJWTMaker(secretKey)
	require.NoError(t, err)
	username := util.RandomOwner()
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token, "token cannot be empty")
	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotZero(t, payload.ID)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
	require.Equal(t, payload.Username, username)
}

func TestExpiredToken(t *testing.T) {
	secretKey := util.RandomString(32)

	maker, err := NewJWTMaker(secretKey)
	require.NoError(t, err)
	require.NotEmpty(t, maker)

	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, errExpiredTokenMessage.Error())
	require.Nil(t, payload)
}

func TestInvalidToken(t *testing.T) {
	payload, err := NewPayload(util.RandomString(6), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)
	secretKey := util.RandomString(32)

	maker, err := NewJWTMaker(secretKey)
	require.NoError(t, err)
	require.NotEmpty(t, maker)
	_, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, errInvalidToken.Error())
}
