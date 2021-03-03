package fakesUserRepository

import (
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
)

type FakeUserTokenRepository struct {
	tokens []entities.UserToken
}

func (r *FakeUserTokenRepository) Generate(userId string) (*entities.UserToken, error) {
	userToken, err := entities.NewUserToken(userId)
	if err != nil{
		return nil, err
	}

	r.tokens = append(r.tokens, *userToken)

	return userToken, nil
}
