package services

import (
	"errors"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
	"ismaeldf/golang-gobarber/modules/users/repositories"
	"os"
	"strings"
)

var FileDirectory = "images/"

type updateUserAvatarService struct {
	usersRepository repositories.IUserRepository
}

func NewUpdateUserAvatarService(repository repositories.IUserRepository) *updateUserAvatarService {
	return &updateUserAvatarService{usersRepository: repository}
}

func (s *updateUserAvatarService) Execute(userId string, filename string) (*entities.User, error) {
	user := s.usersRepository.FindById(userId)
	if user.Id == "" {
		return nil, errors.New("User not exists")
	}

	removeFilePreviousAvatar(user.Avatar)

	user.Avatar = normalizeFilename(filename)

	userUpdated := s.usersRepository.Update(user)

	return &userUpdated, nil
}

func removeFilePreviousAvatar(filename string) {
	_ = os.Remove(FileDirectory + filename)
}

func normalizeFilename(filename string) string {
	return strings.ReplaceAll(filename, FileDirectory, "")
}
