package entities

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewUserToken(t *testing.T) {
	t.Run("should be able to create a new entitie userToken", func(t *testing.T) {
		user, _ := NewUser("jhon doe", "jhondoe@email.com", "12345")

		userToken, _ := NewUserToken(user.Id)

		require.Equal(t, userToken.UserId, user.Id)
	})

	t.Run("should be not able to create a new entitie userToken", func(t *testing.T) {
		_, err := NewUserToken("non-uuid")

		require.NotNil(t, err)
	})
}
