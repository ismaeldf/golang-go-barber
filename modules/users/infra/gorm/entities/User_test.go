package entities

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewUser(t *testing.T) {
	t.Run("should be able to create a new user", func(t *testing.T) {
		user, _ := NewUser("jhon doe", "jhondoe@email.com", "12345")

		require.Equal(t, user.Name, "jhon doe")
	})

	t.Run("should be not able to create a new user", func(t *testing.T) {
		_, err := NewUser("", "", "")

		require.NotNil(t, err)
	})
}
