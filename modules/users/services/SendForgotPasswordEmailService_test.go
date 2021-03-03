package services_test

import (
	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
	fakeHashProvider "ismaeldf/golang-gobarber/modules/users/providers/HashProvider/fakes"
	fakesUserRepository "ismaeldf/golang-gobarber/modules/users/repositories/fakes"
	"ismaeldf/golang-gobarber/modules/users/services"
	"ismaeldf/golang-gobarber/shared/container/providers/MailProvider/fakes"
	"testing"
)

var usersRepository2 fakesUserRepository.FakeUsersRepository
var fakeHashProvider2 fakeHashProvider.FakeHashProvider
var fakeMailProvider2 fakes.FakeMailProvider
var userService *services.CreateUserService
var sendForgotPasswordEMail *services.SendForgotPasswordEmailService
var userTokeRepository2 fakesUserRepository.FakeUserTokenRepository

func beforeEach2(){
	usersRepository2 = fakesUserRepository.FakeUsersRepository{}
	fakeHashProvider2 = fakeHashProvider.FakeHashProvider{}
	fakeMailProvider2 = fakes.FakeMailProvider{}
	userTokeRepository2 = fakesUserRepository.FakeUserTokenRepository{}

	userService = services.NewCreateUserService(&usersRepository2, &fakeHashProvider2)
	sendForgotPasswordEMail = services.NewSendForgotPasswordEmailService(
		&usersRepository2,
		&fakeMailProvider2,
		&userTokeRepository2,
	)

}

func TestSendForgotPasswordEmailService_Execute(t *testing.T) {
	t.Run("should be able to send email forgot password", func(t *testing.T) {
		beforeEach2()

		user := entities.UserUnhide{}
		user.Name = "Jhon Doe"
		user.Email = "jhondoe@email.com"
		user.Password = "12345"

		_, _ = userService.Execute(user)

		err := sendForgotPasswordEMail.Execute(user.Email)

		require.Nil(t, err)
	})

	t.Run("should not be able to send email forgot password using no existing user", func(t *testing.T) {
		beforeEach2()

		err := sendForgotPasswordEMail.Execute("non-user@email.com")

		require.NotNil(t, err)
	})
}
