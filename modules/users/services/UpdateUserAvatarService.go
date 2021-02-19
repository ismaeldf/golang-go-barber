package services

import (
	"errors"
	"ismaeldf/golang-gobarber/config"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
	models1 "ismaeldf/golang-gobarber/modules/users/repositories"
	models2 "ismaeldf/golang-gobarber/shared/container/providers/models"
	"mime/multipart"
	"strings"
)


type updateUserAvatarService struct {
	usersRepository models1.IUserRepository
	storageProvider models2.IStorageProvider
}

func NewUpdateUserAvatarService(repository models1.IUserRepository, storageProvider models2.IStorageProvider) *updateUserAvatarService {
	return &updateUserAvatarService{usersRepository: repository, storageProvider: storageProvider}
}

func (s *updateUserAvatarService) Execute(userId string, file multipart.File) (*entities.User, error) {
	user := s.usersRepository.FindById(userId)
	if user.Id == "" {
		return nil, errors.New("User not exists")
	}

	if user.Avatar != "" {
		s.storageProvider.DeleteFile(user.Avatar)
	}

	filename := s.storageProvider.SaveFile(file)

	user.Avatar = normalizeFilename(filename)

	userUpdated := s.usersRepository.Update(user)

	return &userUpdated, nil
}


func normalizeFilename(filename string) string {
	return strings.ReplaceAll(filename, config.FileDirectory, "")
}
