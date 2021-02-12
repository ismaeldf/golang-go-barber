package entities

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
	"github.com/go-playground/validator/v10"
	"time"
)

type Appointment struct {
	gorm.Model
	Id         string        `json:"id" gorm:"type:uuid;primary_key"`
	ProviderId string        `json:"provider_id" validate:"required"`
	Provider   entities.User `gorm:"foreignKey:ProviderId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Date       time.Time     `json:"date" gorm:"notnull" validate:"required"`
	CreatedAt  time.Time     `json:"create_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
}

func (a *Appointment) isValid() error {
	v := validator.New()
	err := v.Struct(a)
	if err != nil {
		fmt.Errorf("Error during Transaction validation: %s", err.Error())
		return err
	}
	return nil
}

func NewAppointment(providerId string, date time.Time) (*Appointment, error) {
	appointment := Appointment{
		Id:         uuid.NewV4().String(),
		ProviderId: providerId,
		Date:       date,
	}

	err := appointment.isValid()
	if err != nil {
		return nil, err
	}

	return &appointment, nil
}
