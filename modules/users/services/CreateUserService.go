package services

import (
	"errors"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
	"ismaeldf/golang-gobarber/modules/users/repositories"
)

type createUserService struct {
	usersRepository repositories.IUserRepository
}

func NewCreateUserService(repository repositories.IUserRepository) *createUserService {
	return &createUserService{repository}
}

func (s *createUserService) Execute(user entities.UserUnhide) (*entities.User, error) {
	find := s.usersRepository.FindByEmail(user.Email)
	if find.Id != "" {
		return nil, errors.New("This email is already in use")
	}

	userCreated, err := s.usersRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return userCreated, nil
}
