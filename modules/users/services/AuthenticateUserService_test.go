package services_test

import (
	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
	fakeHashProvider "ismaeldf/golang-gobarber/modules/users/providers/HashProvider/fakes"
	fakesUserRepository "ismaeldf/golang-gobarber/modules/users/repositories/fakes"
	"ismaeldf/golang-gobarber/modules/users/services"
	"testing"
)

func TestAuthenticateUserService_Execute(t *testing.T) {
	t.Run("should be able to authenticate user", func(t *testing.T) {
		usersRepository := fakesUserRepository.FakeUsersRepository{}
		fakeHashProvider := fakeHashProvider.FakeHashProvider{}

		userService := services.NewCreateUserService(&usersRepository, &fakeHashProvider)

		user := entities.UserUnhide{}
		user.Name = "Jhon Doe"
		user.Email = "jhondoe@email.com"
		user.Password = "12345"

		_, _ = userService.Execute(user)

		userAuthenticate := services.NewAuthenticateUserService(&usersRepository, &fakeHashProvider)

		auth, _ := userAuthenticate.Execute(user.Email, user.Password)

		require.NotEqual(t, auth.Token, "")
	})

}
