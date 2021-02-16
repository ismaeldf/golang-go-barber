package services

import (
	"errors"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
	providers "ismaeldf/golang-gobarber/modules/users/providers/HashProvider/models"
	"ismaeldf/golang-gobarber/modules/users/repositories"
)

type createUserService struct {
	usersRepository repositories.IUserRepository
	hashProvider providers.IHashProvider
}

func NewCreateUserService(repository repositories.IUserRepository, hashProvider providers.IHashProvider) *createUserService {
	return &createUserService{repository, hashProvider}
}

func (s *createUserService) Execute(user entities.UserUnhide) (*entities.User, error) {
	find := s.usersRepository.FindByEmail(user.Email)
	if find.Id != "" {
		return nil, errors.New("This email is already in use")
	}

	user.Password = s.hashProvider.GenerateHash(user.Password)

	userCreated, err := s.usersRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return userCreated, nil
}
