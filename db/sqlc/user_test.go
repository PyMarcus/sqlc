package db

import (
	"context"
	"testing"

	"github.com/PyMarcus/go_sqlc/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	param := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "hashed",
		FullName:       "userp da silva",
		Email:          util.RandomString(4),
	}
	user, err := testQueries.CreateUser(context.Background(), param)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.Username, param.Username)
	require.Equal(t, user.HashedPassword, param.HashedPassword)
	require.Equal(t, user.FullName, param.FullName)
	require.Equal(t, user.Email, param.Email)
}

func TestGetUser(t *testing.T) {
	param := CreateUserParams{
		Username:       "user1",
		HashedPassword: "hashed",
		FullName:       "user1 da silva",
		Email:          "user@email.com",
	}
	user, err := testQueries.GetUser(context.Background(), param.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.Username, param.Username)
	require.Equal(t, user.HashedPassword, param.HashedPassword)
	require.Equal(t, user.FullName, param.FullName)
	require.Equal(t, user.Email, param.Email)
}
