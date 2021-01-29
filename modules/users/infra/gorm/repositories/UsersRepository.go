package repositories

import (
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
	"ismaeldf/golang-gobarber/shared/infra/gorm"
)

type UsersRepository struct {}

func (r *UsersRepository) Update(user entities.User) entities.User {
	gorm.DB.Model(entities.User{}).Omit("id", "password").Where("id = ?", user.Id).Updates(user)

	return user
}

func (r *UsersRepository) FindById(id string) entities.User {
	var user entities.User

	gorm.DB.Where("id = ?", id).Find(&user)

	return user
}

func (r *UsersRepository) FindByEmail(email string) entities.User {
	var user entities.User

	gorm.DB.Where("email = ?", email).Find(&user)

	return user
}

func (r *UsersRepository) Create(data entities.User) (*entities.User, error){
	user := entities.NewUser(data.Name, data.Email, data.Password)

	err := gorm.DB.Create(user).Error
	if err != nil{
		return nil, err
	}

	return user, nil
}
