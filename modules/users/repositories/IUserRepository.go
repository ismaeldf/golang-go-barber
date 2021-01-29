package repositories

import "ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"

type IUserRepository interface{
	FindById(id string) entities.User
	FindByEmail(email string) entities.User
	Create(data entities.User) (*entities.User, error)
	Update(user entities.User) entities.User
}
