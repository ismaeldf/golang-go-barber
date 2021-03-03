package entities

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type UserToken struct {
	Id        string    `json:"id" gorm:"type:uuid;primary_key" valid:"uuid"`
	UserId      string    `json:"user_id" gorm:"type:uuid" valid:"uuid"`
	Token     string    `json:"token" gorm:"type:uuid,notnull" valid:"uuid"`
	CreatedAt time.Time `json:"create_at" gorm:"autoCreateTime" valid:"-"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" valid:"-"`
}

func (u *UserToken) isValid() error {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}


func NewUserToken(userId string) (*UserToken, error) {
	userToken := UserToken{
		Id:       uuid.NewV4().String(),
		UserId: userId,
		Token: uuid.NewV4().String(),
		CreatedAt: time.Now(),
	}

	err := userToken.isValid()
	if err != nil {
		return nil, err
	}

	return &userToken, nil
}
