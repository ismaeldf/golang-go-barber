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

var usersRepositoryAuthenticateUserService fakesUser.FakeUsersRepository
var fakeHashProviderAuthenticateUserService fakeHashProvider.FakeHashProvider
var fakeTokenProviderAuthenticateUserService fakeTokenProvider.FakeTokenProvider
var userServiceAuthenticateUserService *services.CreateUserService
var userAuthenticate *services.AuthenticateUserService

func beforeEach() {
	usersRepositoryAuthenticateUserService = fakesUser.FakeUsersRepository{}
	fakeHashProviderAuthenticateUserService = fakeHashProvider.FakeHashProvider{}
	fakeTokenProviderAuthenticateUserService = fakeTokenProvider.FakeTokenProvider{}

	userServiceAuthenticateUserService = services.NewCreateUserService(&usersRepositoryAuthenticateUserService, &fakeHashProviderAuthenticateUserService)
	userAuthenticate = services.NewAuthenticateUserService(&usersRepositoryAuthenticateUserService, &fakeHashProviderAuthenticateUserService, &fakeTokenProviderAuthenticateUserService)
}

func TestAuthenticateUserService_Execute(t *testing.T) {
	beforeEach()

	t.Run("should be able to authenticate user", func(t *testing.T) {
		user := entities.UserUnhide{}
		user.Name = "Jhon Doe"
		user.Email = "jhondoe@email.com"
		user.Password = "12345"

		_, _ = userServiceAuthenticateUserService.Execute(user)

		auth, _ := userAuthenticate.Execute(user.Email, user.Password)

		require.NotEqual(t, auth.Token, "")
	})

	t.Run("should be not able to authenticate with non exists user", func(t *testing.T) {
		beforeEach()

		_, err := userAuthenticate.Execute("user-not-exists@email.com", "12345")

		require.NotNil(t, err)
	})

	t.Run("should be not able to authenticate with wrong password", func(t *testing.T) {
		beforeEach()

		user := entities.UserUnhide{}
		user.Name = "Jhon Doe"
		user.Email = "jhondoe@email.com"
		user.Password = "12345"

		_, _ = userServiceAuthenticateUserService.Execute(user)

		_, err := userAuthenticate.Execute(user.Email, "11111")

		require.NotNil(t, err)
	})

}
