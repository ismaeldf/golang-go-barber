package services

import (
	"errors"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
	repositories2 "ismaeldf/golang-gobarber/modules/users/repositories"
	"os"
	"strings"
)

var FileDirectory = "images/"

type updateUserAvatarService struct {
	usersRepository *repositories2.UsersRepository
}

func NewUpdateUserAvatarService(repository *repositories2.UsersRepository) *updateUserAvatarService {
	return &updateUserAvatarService{
		usersRepository: repository,
	}
}

func (s *updateUserAvatarService) Execute(userId string, filename string) (*entities.User, error) {
	user := s.usersRepository.FindById(userId)
	if user.Id == "" {
		return nil, errors.New("User not exists")
	}

	removeFilePreviousAvatar(user.Avatar)

	user.Avatar = normalizeFilename(filename)

	userUpdated := s.usersRepository.UpdateAvatar(user)

	return &userUpdated, nil
}

func removeFilePreviousAvatar(filename string) {
	_ = os.Remove(FileDirectory + filename)
}

func normalizeFilename(filename string) string {
	return strings.ReplaceAll(filename, FileDirectory, "")
}
