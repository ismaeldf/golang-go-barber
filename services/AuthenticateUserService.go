package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"ismaeldf.melo/golang/go-barber/models"
	"ismaeldf.melo/golang/go-barber/repositories"
)

type authenticateUserService struct {
	usersRepository *repositories.UsersRepository
}

const errorMsg = "Incorrect Email/Password combination"

func NewAuthenticateUserService(repository *repositories.UsersRepository) *authenticateUserService {
	return &authenticateUserService{repository}
}

func (s *authenticateUserService) Execute(email string, password string) (*models.User, error) {
	user := s.usersRepository.FindByEmail(email)
	if user.ID == "" {
		return nil, errors.New(errorMsg)
	}

	passwordMatched := isCorrectPassword(user, password)
	if !passwordMatched {
		return nil, errors.New(errorMsg)
	}

	return &user, nil
}

func isCorrectPassword(user models.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}