package services_test

import (
	"github.com/stretchr/testify/require"
	_ "github.com/stretchr/testify/require"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
	fakeHashProvider "ismaeldf/golang-gobarber/modules/users/providers/HashProvider/fakes"
	fakesUserRepository "ismaeldf/golang-gobarber/modules/users/repositories/fakes"
	"ismaeldf/golang-gobarber/modules/users/services"
	"ismaeldf/golang-gobarber/shared/container/providers/StorageProvider/fakes"
	"mime/multipart"
	"testing"
)

func TestUpdateUserAvatarService_Execute(t *testing.T) {
	t.Run("should be able to update new avatar", func(t *testing.T) {
		usersRepository := fakesUserRepository.FakeUsersRepository{}
		fakeHashProvider := fakeHashProvider.FakeHashProvider{}
		fakeStorageProvider := fakes.FakeStorageProvider{}

		userService := services.NewCreateUserService(&usersRepository, &fakeHashProvider)

		user := entities.UserUnhide{}
		user.Name = "Jhon Doe"
		user.Email = "jhondoe@email.com"
		user.Password = "12345"

		userCreated, _ := userService.Execute(user)

		updateUserAvatarService := services.NewUpdateUserAvatarService(
			&usersRepository,
			&fakeStorageProvider,
		)

		var file  multipart.File

		userUpdated, _ := updateUserAvatarService.Execute(userCreated.Id, file)

		require.NotEqual(t, userUpdated.Avatar, "")
	})

	t.Run("should be not able to update avatar when user non exists", func(t *testing.T) {
		usersRepository := fakesUserRepository.FakeUsersRepository{}
		fakeStorageProvider := fakes.FakeStorageProvider{}

		updateUserAvatarService := services.NewUpdateUserAvatarService(
			&usersRepository,
			&fakeStorageProvider,
		)

		var file  multipart.File

		_, err := updateUserAvatarService.Execute("no-user-exists", file)

		require.NotNil(t, err)
	})

	t.Run("should be able to update exists previous avatar", func(t *testing.T) {
		usersRepository := fakesUserRepository.FakeUsersRepository{}
		fakeHashProvider := fakeHashProvider.FakeHashProvider{}
		fakeStorageProvider := fakes.FakeStorageProvider{}

		userService := services.NewCreateUserService(&usersRepository, &fakeHashProvider)

		user := entities.UserUnhide{}
		user.Name = "Jhon Doe"
		user.Email = "jhondoe@email.com"
		user.Password = "12345"

		userCreated, _ := userService.Execute(user)

		updateUserAvatarService := services.NewUpdateUserAvatarService(
			&usersRepository,
			&fakeStorageProvider,
		)

		var file  multipart.File

		userUpdated1, _ := updateUserAvatarService.Execute(userCreated.Id, file)
		userUpdated2, _ := updateUserAvatarService.Execute(userCreated.Id, file)

		require.NotEqual(t, userUpdated1.Avatar, userUpdated2.Avatar)
	})
}
