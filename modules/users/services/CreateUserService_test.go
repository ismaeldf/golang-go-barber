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

func TestCreateUserService_Execute(t *testing.T) {
	t.Run("should be able to create a new user", func(t *testing.T) {
		usersRepository := fakesUserRepository.FakeUsersRepository{}
		fakeHashProvider := fakeHashProvider.FakeHashProvider{}

		userService := services.NewCreateUserService(&usersRepository, &fakeHashProvider)

		user := entities.UserUnhide{}
		user.Name = "Jhon Doe"
		user.Email = "jhondoe@email.com"
		user.Password = "12345"

		userCreated, _ := userService.Execute(user)

		require.Equal(t, userCreated.Name, user.Name)
	})

	t.Run("should be not able to create a new user with same email registered", func(t *testing.T) {
		usersRepository := fakesUserRepository.FakeUsersRepository{}
		fakeHashProvider := fakeHashProvider.FakeHashProvider{}

		userService := services.NewCreateUserService(&usersRepository, &fakeHashProvider)

		user := entities.UserUnhide{}
		user.Name = "Jhon Doe"
		user.Email = "jhondoe@email.com"
		user.Password = "12345"

		_, _ = userService.Execute(user)

		_, err := userService.Execute(user)

		require.NotNil(t, err)
	})

	t.Run("should be not able to create a new user with not valid email", func(t *testing.T) {
		usersRepository := fakesUserRepository.FakeUsersRepository{}
		fakeHashProvider := fakeHashProvider.FakeHashProvider{}

		userService := services.NewCreateUserService(&usersRepository, &fakeHashProvider)

		user := entities.UserUnhide{}
		user.Name = "Jhon Doe"
		user.Email = ""
		user.Password = "12345"

		_, err := userService.Execute(user)

		require.NotNil(t, err)
	})
}
