package repositories

import (
	"ismaeldf/golang-gobarber/database"
	"ismaeldf/golang-gobarber/models"
)

type UsersRepository struct {}

func (r *UsersRepository) UpdateAvatar(user models.User) models.User {
	database.DB.Model(models.User{}).Where("id = ?", user.Id).Update("Avatar", user.Avatar)

	return user
}

func (r *UsersRepository) FindById(id string) models.User {
	var user models.User

	database.DB.Where("id = ?", id).Find(&user)

	return user
}

func (r *UsersRepository) FindByEmail(email string) models.User {
	var user models.User

	database.DB.Where("email = ?", email).Find(&user)

	return user
}

func (r *UsersRepository) Create(data models.User) (*models.User, error){
	user := models.NewUser(data.Name, data.Email, data.Password)

	err := database.DB.Create(user).Error
	if err != nil{
		return nil, err
	}

	return user, nil
}
