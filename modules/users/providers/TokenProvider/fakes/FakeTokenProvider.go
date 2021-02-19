package fakeTokenProvider


type FakeTokenProvider struct{}

func (j *FakeTokenProvider) CreateToken(userId string) string {
	return userId
}

func (j *FakeTokenProvider) DecodeToken(tokenString string) (*string, error){
	return &tokenString, nil
}
