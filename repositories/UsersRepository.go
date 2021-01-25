package repositories

import (
	"ismaeldf.melo/golang/go-barber/database"
	"ismaeldf.melo/golang/go-barber/models"
)

type UsersRepository struct {}

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
