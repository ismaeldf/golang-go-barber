package models

type IMailProvider interface {
	SendMail(to string, body string) error
}
