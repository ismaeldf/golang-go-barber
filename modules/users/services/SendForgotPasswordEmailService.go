package services

import (
	"errors"
	models1 "ismaeldf/golang-gobarber/modules/users/repositories"
	models2 "ismaeldf/golang-gobarber/shared/container/providers/MailProvider/models"
)

type SendForgotPasswordEmailService struct {
	usersRepository     models1.IUserRepository
	mailProvider        models2.IMailProvider
	userTokenRepository models1.IUserTokenRepository
}

func NewSendForgotPasswordEmailService(
	userRepository models1.IUserRepository,
	mailProvider models2.IMailProvider,
	userTokeRepository models1.IUserTokenRepository,
) *SendForgotPasswordEmailService {
	return &SendForgotPasswordEmailService{
		usersRepository:     userRepository,
		mailProvider:        mailProvider,
		userTokenRepository: userTokeRepository,
	}
}

func (s *SendForgotPasswordEmailService) Execute(email string) error {
	user := s.usersRepository.FindByEmail(email)
	if user.Id == "" {
		return errors.New("User does not exists")
	}

	err := s.mailProvider.SendMail(email, "Email de recuperação de senha")
	if err != nil {
		return err
	}
	return nil
}
