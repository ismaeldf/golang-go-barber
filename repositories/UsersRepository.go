package repositories

import (
	"gorm.io/gorm"
	"ismaeldf.melo/golang/go-barber/models"
)

type UsersRepositoryDTO struct {
	Name string
	Email string
	Password string
}

type UsersRepository struct {
	DB *gorm.DB
}

func (r *UsersRepository) FindByEmail(email string) models.User {
	var user models.User
	r.DB.Where("email = ?", email).Find(&user)

	return user
}

func (r *UsersRepository) Create(data UsersRepositoryDTO) (*models.User, error){
	user := models.NewUser(data.Name, data.Email, data.Password)

	err := r.DB.Create(user).Error
	if err != nil{
		return nil, err
	}

	return user, nil
}
