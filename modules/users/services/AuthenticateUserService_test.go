package services_test

import (
	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
	fakeHashProvider "ismaeldf/golang-gobarber/modules/users/providers/HashProvider/fakes"
	fakeTokenProvider "ismaeldf/golang-gobarber/modules/users/providers/TokenProvider/fakes"
	fakesUser "ismaeldf/golang-gobarber/modules/users/repositories/fakes"
	"ismaeldf/golang-gobarber/modules/users/services"
	"testing"
)

var usersRepository1 fakesUser.FakeUsersRepository
var fakeHashProvider1 fakeHashProvider.FakeHashProvider
var fakeTokenProvider1 fakeTokenProvider.FakeTokenProvider
var userService1 *services.CreateUserService
var userAuthenticate *services.AuthenticateUserService

func beforeEach1() {
	usersRepository1 = fakesUser.FakeUsersRepository{}
	fakeHashProvider1 = fakeHashProvider.FakeHashProvider{}
	fakeTokenProvider1 = fakeTokenProvider.FakeTokenProvider{}

	userService1 = services.NewCreateUserService(&usersRepository1, &fakeHashProvider1)
	userAuthenticate = services.NewAuthenticateUserService(&usersRepository1, &fakeHashProvider1, &fakeTokenProvider1)
}

func TestAuthenticateUserService_Execute(t *testing.T) {
	beforeEach1()

	t.Run("should be able to authenticate user", func(t *testing.T) {
		user := entities.UserUnhide{}
		user.Name = "Jhon Doe"
		user.Email = "jhondoe@email.com"
		user.Password = "12345"

		_, _ = userService1.Execute(user)

		auth, _ := userAuthenticate.Execute(user.Email, user.Password)

		require.NotEqual(t, auth.Token, "")
	})

	t.Run("should be not able to authenticate with non exists user", func(t *testing.T) {
		beforeEach1()

		_, err := userAuthenticate.Execute("user-not-exists@email.com", "12345")

		require.NotNil(t, err)
	})

	t.Run("should be not able to authenticate with wrong password", func(t *testing.T) {
		beforeEach1()

		user := entities.UserUnhide{}
		user.Name = "Jhon Doe"
		user.Email = "jhondoe@email.com"
		user.Password = "12345"

		_, _ = userService1.Execute(user)

		_, err := userAuthenticate.Execute(user.Email, "11111")

		require.NotNil(t, err)
	})

}
