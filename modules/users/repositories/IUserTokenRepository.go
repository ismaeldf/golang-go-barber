package repositories

import "ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"

type IUserTokenRepository interface {
	Generate(userId string) (*entities.UserToken, error)
}
