package services

import (
	"errors"
	"ismaeldf.melo/golang/go-barber/models"
	"ismaeldf.melo/golang/go-barber/repositories"
)

type RequestCreateUser struct {
	Name     string
	Email    string
	Password string
}

type createUserService struct {
	usersRepository *repositories.UsersRepository
}

func NewCreateUserService(repository *repositories.UsersRepository) *createUserService {
	return &createUserService{repository}
}

func (s *createUserService) Execute(data RequestCreateUser) (*models.User, error) {
	find := s.usersRepository.FindByEmail(data.Email)
	if find.ID != "" {
		return nil, errors.New("This email is already in use")
	}

	user, err := s.usersRepository.Create(repositories.UsersRepositoryDTO{
		Name: data.Name,
		Email: data.Email,
		Password: data.Password,
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
