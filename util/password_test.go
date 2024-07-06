package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	word := "Secret"
	hashed, err := HashPassword(word)
	require.NoError(t, err)
	require.NotEmpty(t, hashed)
}

func TestCheckPassword(t *testing.T) {
	word := "Secret"
	hashed, err := HashPassword(word)
	require.NoError(t, err)
	require.NotEmpty(t, hashed)
	err = CheckPassword(word, hashed)
	require.NoError(t, err)

	incorrectWord := "OK"
	err = CheckPassword(incorrectWord, hashed)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
