package entities

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"ismaeldf/golang-gobarber/modules/users/infra/gorm/entities"
	"time"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Appointment struct {
	Id         string        `json:"id" gorm:"type:uuid;primary_key" valid:"-"`
	ProviderId string        `json:"provider_id" valid:"notnull"`
	Provider   entities.User `gorm:"foreignKey:ProviderId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" valid:"-"`
	Date       time.Time     `json:"date" gorm:"notnull" valid:"-"`
	CreatedAt  time.Time     `json:"create_at" gorm:"autoCreateTime" valid:"-"`
	UpdatedAt  time.Time     `json:"updated_at" gorm:"autoUpdateTime" valid:"-"`
}

func (a *Appointment) isValid() error {
	_, err := govalidator.ValidateStruct(a)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func NewAppointment(providerId string, date time.Time) (*Appointment, error) {
	appointment := Appointment{
		ProviderId: providerId,
		Date:       date,
	}

	appointment.Id = uuid.NewV4().String()
	appointment.CreatedAt = time.Now()

	err := appointment.isValid()
	if err != nil {
		return nil, err
	}

	return &appointment, nil
}
