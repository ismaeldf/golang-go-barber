package services_test

import (
	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
	fakesUserRepository "ismaeldf/golang-gobarber/modules/users/repositories/fakes"
	"ismaeldf/golang-gobarber/modules/users/services"
	"testing"
)

func TestCreateUserService_Execute(t *testing.T) {
	t.Run("should be able to create a new user", func(t *testing.T) {
		usersRepository := fakesUserRepository.FakeUsersRepository{}

		userService := services.NewCreateUserService(&usersRepository)

		user := entities.UserUnhide{}
		user.Name = "Jhon Doe"
		user.Email = "jhondoe@email.com"
		user.Password = "12345"

		userCreated, _ := userService.Execute(user)

		require.Equal(t, userCreated.Name, user.Name)
	})

}
