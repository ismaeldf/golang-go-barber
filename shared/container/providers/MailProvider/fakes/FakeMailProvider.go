package fakes

type message struct {
	to string
	body string
}

type FakeMailProvider struct {
	Messages []message
}

func (f *FakeMailProvider) SendMail(to string, body string) error{
	message := message{to, body}
	f.Messages = append(f.Messages, message)
	return nil
}
