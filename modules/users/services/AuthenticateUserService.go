package services

import (
	"errors"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
	model1 "ismaeldf/golang-gobarber/modules/users/providers/HashProvider/models"
	model2 "ismaeldf/golang-gobarber/modules/users/providers/TokenProvider/models"
	"ismaeldf/golang-gobarber/modules/users/repositories"
)

type ResponseAuthenticateUser struct {
	User  entities.User
	Token string
}

type authenticateUserService struct {
	usersRepository repositories.IUserRepository
	hashProvider    model1.IHashProvider
	tokenProvider   model2.ITokenProvider
}

const errorMsg = "Incorrect Email/Password combination"

func NewAuthenticateUserService(
	repository repositories.IUserRepository,
	hashProvider model1.IHashProvider,
	tokenProvider model2.ITokenProvider,
) *authenticateUserService {
	return &authenticateUserService{
		usersRepository: repository,
		hashProvider: hashProvider,
		tokenProvider: tokenProvider,
	}
}

func (s *authenticateUserService) Execute(email string, password string) (*ResponseAuthenticateUser, error) {
	user := s.usersRepository.FindByEmail(email)
	if user.Id == "" {
		return nil, errors.New(errorMsg)
	}

	passwordMatched := s.isCorrectPassword(user, password)
	if !passwordMatched {
		return nil, errors.New(errorMsg)
	}

	token := s.tokenProvider.CreateToken(user.Id)

	response := ResponseAuthenticateUser{
		User:  user,
		Token: token,
	}

	return &response, nil
}

func (a *authenticateUserService) isCorrectPassword(user entities.User, password string) bool {
	return a.hashProvider.CompareHash(password, user.Password)
}
