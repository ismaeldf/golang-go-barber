package services

import (
	"errors"
	"ismaeldf/golang-gobarber/models"
	"ismaeldf/golang-gobarber/repositories"
)

type createUserService struct {
	usersRepository *repositories.UsersRepository
}

func NewCreateUserService(repository *repositories.UsersRepository) *createUserService {
	return &createUserService{repository}
}

func (s *createUserService) Execute(user models.User) (*models.User, error) {
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
