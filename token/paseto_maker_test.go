package token

import (
	"testing"
	"time"

	"github.com/PyMarcus/go_sqlc/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T){
	simetricKey := util.RandomString(32)
	maker, err := NewPasetoMaker(simetricKey)
	require.NoError(t, err)
	require.NotEmpty(t, maker)
	userName := util.RandomOwner()
	token, err := maker.CreateToken(userName, time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.Equal(t, payload.Username, userName)
	require.NotZero(t, payload.ID)
}