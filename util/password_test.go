package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(wrongPassword))
	require.Error(t, err, bcrypt.ErrMismatchedHashAndPassword)
}
