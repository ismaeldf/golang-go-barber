package fakesUserRepository

import (
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
)

type FakeUsersRepository struct {
	users []entities.User
}

func (r *FakeUsersRepository) Update(user entities.User) entities.User {
	for i, u := range r.users {
		if u.Id == user.Id {
			r.users[i] = user
			return r.users[i]
		}
	}

	return user
}

func (r *FakeUsersRepository) FindById(id string) entities.User {
	var user entities.User

	for _, u := range r.users {
		if u.Id == id {
			return u
		}
	}

	return user
}

func (r *FakeUsersRepository) FindByEmail(email string) entities.User {
	var user entities.User

	for _, u := range r.users {
		if u.Email == email {
			return u
		}
	}

	return user
}

func (r *FakeUsersRepository) Create(data entities.UserUnhide) (*entities.User, error){
	user, err := entities.NewUser(data.Name, data.Email, data.Password)
	if err != nil{
		return nil, err
	}

	r.users = append(r.users, *user)

	return user, nil
}
