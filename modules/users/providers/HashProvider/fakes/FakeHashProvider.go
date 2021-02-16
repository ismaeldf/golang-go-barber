package fakeHashProvider

type FakeHashProvider struct {}

func (b *FakeHashProvider) GenerateHash(payload string) string {
	return payload
}


func (b *FakeHashProvider) CompareHash(payload string, hashed string) bool {
	return payload == hashed
}
