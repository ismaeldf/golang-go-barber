package services

import (
	models1 "ismaeldf/golang-gobarber/modules/users/repositories"
	models2 "ismaeldf/golang-gobarber/shared/container/providers/MailProvider/models"
)


type sendForgotPasswordEmailService struct {
	usersRepository models1.IUserRepository
	mailProvider models2.IMailProvider
}

func NewSendForgotPasswordEmailService(repository models1.IUserRepository, mailProvider models2.IMailProvider) *sendForgotPasswordEmailService {
	return &sendForgotPasswordEmailService{usersRepository: repository, mailProvider: mailProvider}
}

func (s *sendForgotPasswordEmailService) Execute(email string) error{
	err := s.mailProvider.SendMail(email, "Email de recuperação de senha")
	if err != nil {
		return err
	}
	return nil
}

