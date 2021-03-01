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

func TestSendForgotPasswordEmailService_Execute(t *testing.T) {
	t.Run("should be able to send email forgot password", func(t *testing.T) {
		usersRepository := fakesUserRepository.FakeUsersRepository{}
		fakeHashProvider := fakeHashProvider.FakeHashProvider{}
		fakeMailProvider := fakes.FakeMailProvider{}

		userService := services.NewCreateUserService(&usersRepository, &fakeHashProvider)

		user := entities.UserUnhide{}
		user.Name = "Jhon Doe"
		user.Email = "jhondoe@email.com"
		user.Password = "12345"

		_, _ = userService.Execute(user)

		sendForgotPasswordEMail := services.NewSendForgotPasswordEmailService(&usersRepository, &fakeMailProvider)

		err := sendForgotPasswordEMail.Execute(user.Email)

		require.Nil(t, err)
	})

	t.Run("should not be able to send email forgot password using no existing user", func(t *testing.T) {
		usersRepository := fakesUserRepository.FakeUsersRepository{}
		fakeMailProvider := fakes.FakeMailProvider{}

		sendForgotPasswordEMail := services.NewSendForgotPasswordEmailService(&usersRepository, &fakeMailProvider)

		err := sendForgotPasswordEMail.Execute("non-user@email.com")

		require.NotNil(t, err)
	})
}
